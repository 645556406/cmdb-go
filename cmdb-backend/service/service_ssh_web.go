package service

import (
	"cmdb-backend/dao"
	"cmdb-backend/utils"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // 允许跨域（生产环境需限制）
}

// 自定义Writer将SSH输出转发到WebSocket
type sshWriter struct{ ws *websocket.Conn }

func (w *sshWriter) Write(p []byte) (n int, err error) {
	err = w.ws.WriteMessage(websocket.TextMessage, p)
	return len(p), err
}

// HandleWebSSHSinger 处理WebSocket SSH连接请求，将HTTP连接升级为WebSocket连接，用于实现Web SSH功能, 使用私钥进行认证。
// 参数：
// c *gin.Context: gin框架的上下文对象，用于处理HTTP请求和响应。
func HandleWebSSHSinger(c *gin.Context) {
	log.Println("收到连接请求")
	// 1. 升级HTTP连接为WebSocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Println("WebSocket关闭失败:", err)
			return
		}
	}(ws)

	// 2. 读取前端发送的连接参数
	var params struct {
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ws.ReadJSON(&params); err != nil {
		log.Println("参数解析失败:", err)
		return
	}
	// 加载配置文件，获取私钥路径
	configYaml := utils.LoadYamlConfigNew("config/dev.yaml")
	privateKeyPath := configYaml["SSHPrivateKey"].(map[string]interface{})["path"].(string)
	if privateKeyPath == "" {
		log.Println("私钥路径为空")
		return
	}
	log.Println("私钥路径:", privateKeyPath)
	// 3. 读取并解析私钥文件
	privateKeyBytes, errReadFile := os.ReadFile(privateKeyPath)
	if errReadFile != nil {
		return
	}

	// 4. 解析私钥（支持加密的私钥）
	signer, errSSH := ssh.ParsePrivateKey(privateKeyBytes)
	if errSSH != nil {
		// 如果私钥有密码，尝试解密
		var passphraseMissingError *ssh.PassphraseMissingError
		if errors.As(errSSH, &passphraseMissingError) {
			// 尝试使用密码解密私钥(如果有的话这里需要修改下ssh私钥密码，没有就可以忽略了)
			signer, err = decryptPrivateKey(privateKeyBytes, "")
			if err != nil {
				return
			}
		}
	}
	// 获取主机公钥
	host := params.Host
	log.Println("主机:", host)
	serverInfo, err := dao.GetServerOneByIP(host)
	if err != nil {
		log.Println(err)
		return
	}
	publicKey := serverInfo.PublicKey
	if publicKey == "" {
		log.Println("主机公钥为空")
		return
	}
	// 从文件或硬编码字符串加载公钥
	hostKeyBytes := []byte(publicKey)
	hostKey, _, _, _, err := ssh.ParseAuthorizedKey(hostKeyBytes)
	if err != nil {
		log.Fatal("解析公钥失败:", err)
	}
	// 5. 配置 SSH 客户端
	config := &ssh.ClientConfig{
		User:              params.Username,
		HostKeyAlgorithms: []string{"ssh-rsa", "rsa-sha2-256"}, // 指定主机密钥算法,不然可以出现解密失败的问题
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer), // 使用私钥认证
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey), // 生产环境应改为 ssh.FixedHostKey // 开发环境改为ssh.InsecureIgnoreHostKey()
		Timeout:         10 * time.Second,
	}

	// 6.建立 SSH 连接
	sshClient, errC := ssh.Dial("tcp", params.Host+":22", config)
	if errC != nil {
		err := ws.WriteJSON(map[string]string{"error": "SSH连接失败: " + errC.Error()})
		if err != nil {
			return
		}
		return
	}
	defer func(sshClient *ssh.Client) {
		err := sshClient.Close()
		if err != nil {

		}
	}(sshClient)

	// 创建SSH会话
	session, err := sshClient.NewSession()
	if err != nil {
		err := ws.WriteJSON(map[string]string{"error": "创建会话失败"})
		if err != nil {
			return
		}
		return
	}
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {
			return
		}
	}(session)

	// 7. 绑定标准输入输出
	session.Stdout = &sshWriter{ws: ws}
	session.Stderr = &sshWriter{ws: ws}
	stdinPipe, _ := session.StdinPipe()

	// 8. 设置终端参数并启动 Shell
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 启用回显
		ssh.TTY_OP_ISPEED: 14400, // 输入速度
		ssh.TTY_OP_OSPEED: 14400, // 输出速度
	}
	if err := session.RequestPty("xterm", 55, 180, modes); err != nil {
		log.Println("PTY请求失败:", err)
		return
	}

	// 启动Shell
	if err := session.Shell(); err != nil {
		log.Println("启动Shell失败:", err)
		return
	}

	// 9. 转发 WebSocket 消息到 SSH
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// 如果是 resize 消息，不转发到 SSH
		if strings.HasPrefix(string(msg), `{"type":"resize"`) {
			var resizeMsg struct{ Cols, Rows int }
			err := json.Unmarshal(msg, &resizeMsg)
			if err != nil {
				return
			}
			err = session.WindowChange(resizeMsg.Rows, resizeMsg.Cols)
			if err != nil {
				return
			} // 调整 PTY 尺寸
			continue // 跳过后续转发逻辑
		}
		_, errStdin := stdinPipe.Write(msg)
		if errStdin != nil {
			return
		}
	}
}

// 解密加密的私钥
func decryptPrivateKey(keyBytes []byte, passphrase string) (ssh.Signer, error) {
	return ssh.ParsePrivateKeyWithPassphrase(keyBytes, []byte(passphrase))
}

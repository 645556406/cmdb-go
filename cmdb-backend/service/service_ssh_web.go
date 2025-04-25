package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // 允许跨域（生产环境需限制）
}

// HandleWebSSH 处理WebSocket SSH连接请求, 使用密码进行鉴权
func HandleWebSSH(c *gin.Context) {
	log.Println("收到连接请求")
	// 升级HTTP连接为WebSocket
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

	// 读取前端发送的连接参数
	var params struct {
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ws.ReadJSON(&params); err != nil {
		log.Println("参数解析失败:", err)
		return
	}

	// 建立SSH连接
	config := &ssh.ClientConfig{
		User: params.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(params.Password), // 密码认证
			//ssh.PublicKeys(signer), // 使用私钥认证
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 生产环境应验证HostKey
		Timeout:         10 * time.Second,
	}
	log.Println(params)
	sshClient, errC := ssh.Dial("tcp", params.Host+":22", config)
	if errC != nil {
		ws.WriteJSON(map[string]string{"error": "SSH连接失败: " + errC.Error()})
		return
	}
	defer sshClient.Close()

	// 创建SSH会话
	session, err := sshClient.NewSession()
	if err != nil {
		ws.WriteJSON(map[string]string{"error": "创建会话失败"})
		return
	}
	defer session.Close()

	// 绑定标准输入输出
	session.Stdout = &sshWriter{ws: ws}
	session.Stderr = &sshWriter{ws: ws}
	stdinPipe, _ := session.StdinPipe()

	// 设置终端参数
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 启用回显
		ssh.TTY_OP_ISPEED: 14400, // 输入速度
		ssh.TTY_OP_OSPEED: 14400, // 输出速度
	}
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Println("PTY请求失败:", err)
		return
	}

	// 启动Shell
	if err := session.Shell(); err != nil {
		log.Println("启动Shell失败:", err)
		return
	}

	// 转发WebSocket消息到SSH
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// 如果是 resize 消息，不转发到 SSH
		if strings.HasPrefix(string(msg), `{"type":"resize"`) {
			var resizeMsg struct{ Cols, Rows int }
			json.Unmarshal(msg, &resizeMsg)
			session.WindowChange(resizeMsg.Rows, resizeMsg.Cols) // 调整 PTY 尺寸
			continue                                             // 跳过后续转发逻辑
		}
		stdinPipe.Write(msg)
	}
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
	// 升级HTTP连接为WebSocket
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

	// 读取前端发送的连接参数
	var params struct {
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ws.ReadJSON(&params); err != nil {
		log.Println("参数解析失败:", err)
		return
	}
	// 1. 读取私钥文件
	privateKeyBytes, err := ioutil.ReadFile("/Users/wangye/.ssh/id_rsa")
	fmt.Println(string(privateKeyBytes))
	if err != nil {
		return
	}

	// 2. 解析私钥（支持加密的私钥）
	signer, errSSH := ssh.ParsePrivateKey(privateKeyBytes)
	if errSSH != nil {
		// 如果私钥有密码，尝试解密
		if _, ok := errSSH.(*ssh.PassphraseMissingError); ok {
			// 尝试使用密码解密私钥(如果有的话这里需要修改下ssh私钥密码，没有就可以忽略了)
			signer, err = decryptPrivateKey(privateKeyBytes, "your-passphrase")
			if err != nil {
				return
			}
		} else {
			return
		}
	}

	// 3. 配置 SSH 客户端
	config := &ssh.ClientConfig{
		User: params.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer), // 使用私钥认证
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 生产环境应改为 ssh.FixedHostKey
		Timeout:         10 * time.Second,
	}

	// 4. 建立连接
	sshClient, errC := ssh.Dial("tcp", params.Host+":22", config)
	if errC != nil {
		ws.WriteJSON(map[string]string{"error": "SSH连接失败: " + errC.Error()})
		return
	}
	defer sshClient.Close()

	// 创建SSH会话
	session, err := sshClient.NewSession()
	if err != nil {
		ws.WriteJSON(map[string]string{"error": "创建会话失败"})
		return
	}
	defer session.Close()

	// 绑定标准输入输出
	session.Stdout = &sshWriter{ws: ws}
	session.Stderr = &sshWriter{ws: ws}
	stdinPipe, _ := session.StdinPipe()

	// 设置终端参数
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 启用回显
		ssh.TTY_OP_ISPEED: 14400, // 输入速度
		ssh.TTY_OP_OSPEED: 14400, // 输出速度
	}
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Println("PTY请求失败:", err)
		return
	}

	// 启动Shell
	if err := session.Shell(); err != nil {
		log.Println("启动Shell失败:", err)
		return
	}

	// 转发WebSocket消息到SSH
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// 如果是 resize 消息，不转发到 SSH
		if strings.HasPrefix(string(msg), `{"type":"resize"`) {
			var resizeMsg struct{ Cols, Rows int }
			json.Unmarshal(msg, &resizeMsg)
			session.WindowChange(resizeMsg.Rows, resizeMsg.Cols) // 调整 PTY 尺寸
			continue                                             // 跳过后续转发逻辑
		}
		stdinPipe.Write(msg)
	}
}

// 解密加密的私钥
func decryptPrivateKey(keyBytes []byte, passphrase string) (ssh.Signer, error) {
	return ssh.ParsePrivateKeyWithPassphrase(keyBytes, []byte(passphrase))
}

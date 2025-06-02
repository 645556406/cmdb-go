package service

import (
	"cmdb-backend/dao"
	"cmdb-backend/model"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ping/ping"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// GetServerList 从 DAO 层获取服务器列表并返回给客户端
//
// 参数:
//
//	c: *gin.Context - Gin 框架的上下文对象，用于处理 HTTP 请求和响应
//
// 返回值:
//
//	无
func GetServerList(c *gin.Context) {
	serverList := dao.GetServerList()
	var response Response
	response.Code = 20000
	response.Message = "success"
	response.Data = serverList
	c.JSON(http.StatusOK, response)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// AddServer 函数用于处理添加服务器的请求
//
// 参数:
//
//	context: *gin.Context - HTTP请求的上下文对象
//
// 返回值:
//
//	无
func AddServer(context *gin.Context) {
	var server model.Server
	if err := context.ShouldBindJSON(&server); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := dao.AddServer(server)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, Response{Code: 20000, Message: "success"})
	}
}

// DelServer 用于删除服务器信息
//
// 参数:
//
//	context: gin的上下文对象，用于处理HTTP请求
//
// 返回值:
//
//	无
func DelServer(context *gin.Context) {
	var server model.Server
	if err := context.ShouldBindJSON(&server); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := dao.DelServer(server)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, Response{Code: 20000, Message: "success"})
	}
}

// UpdateServer 用于更新服务器信息
//
// 参数:
//
//	context: *gin.Context - gin框架的上下文对象
//
// 返回值:
//
//	无
func UpdateServer(context *gin.Context) {
	var server model.Server
	if err := context.ShouldBindJSON(&server); err != nil {
		// 继续执行后续逻辑
		log.Println("err:", err)
		context.JSON(http.StatusBadRequest, Response{Code: 40000, Message: err.Error(), Data: server})
		return
	}
	log.Println("server:", server)
	err := dao.UpdateServer(server)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, Response{Code: 20000, Message: "success"})
	}
}

// GetServerDetailByID 函数通过服务器ID获取服务器详细信息
//
// 参数：
//
//	context *gin.Context：Gin框架的上下文对象，用于处理HTTP请求
//
// 返回值：
//
//	无
func GetServerDetailByID(context *gin.Context) {
	serverID := context.Param("id")
	log.Println("GetServerDetailByID:", serverID)
	serverIDUint, err := strconv.ParseUint(serverID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	serverFirst, errDetail := dao.GetServerDetailByID(serverIDUint)
	if errDetail != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": errDetail.Error()})
		return
	} else {
		context.JSON(http.StatusOK, Response{Code: 20000, Message: "success", Data: serverFirst})
	}
}

// GetServerOneByIP 是一个根据IP获取服务器信息的处理函数
func GetServerOneByIP(context *gin.Context) {
	IP := context.Param("ip")
	ipv4Pattern := `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	re := regexp.MustCompile(ipv4Pattern)
	if !re.MatchString(IP) {
		context.JSON(http.StatusOK, Response{Code: 40001, Message: "IP参数格式错误"})
		return
	}
	if IP == "" {
		context.JSON(http.StatusOK, Response{Code: 40001, Message: "IP参数不能为空"})
		return
	}
	server, err := dao.GetServerOneByIP(IP)
	if err != nil {
		context.JSON(http.StatusOK, Response{Code: 50000, Message: "IP不存在", Data: server})
		return
	}
	context.JSON(http.StatusOK, Response{Code: 20000, Message: "success", Data: server})
}

// GetCountServer 函数用于获取服务器总数
func GetCountServer(context *gin.Context) {
	count, err := dao.GetServerCount()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, Response{Code: 20000, Message: "success", Data: count})
	}
}

// GetOnlineCountServer 函数用于获取在线服务器的数量
func GetOnlineCountServer(context *gin.Context) {
	count, err := dao.GetOnlineCountServer()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, Response{Code: 20000, Message: "success", Data: count})
	}
}

// GetOfflineCountServer 函数用于获取离线服务器的数量
func GetOfflineCountServer(context *gin.Context) {
	count, err := dao.GetOfflineCountServer()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, Response{Code: 20000, Message: "success", Data: count})
	}
}

func GetServerIPList() ([]model.Server, error) {
	serverList, err := dao.GetServerIPList()
	return serverList, err
}

func CheckServerStatus(ip string) bool {
	pingEr, err := ping.NewPinger(ip)
	if err != nil {
		return false
	}
	pingEr.SetPrivileged(true)
	pingEr.Count = 2
	pingEr.Timeout = 2 * time.Second
	err = pingEr.Run()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	stats := pingEr.Statistics()
	return stats.PacketsRecv > 0
}

func MacPing(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// UpdateServerStatus 函数用于更新服务器的状态。
func UpdateServerStatus(server model.Server) {
	var result bool
	switch runtime.GOOS {
	case "darwin":

		result = MacPing(server.IP)
	default:
		result = CheckServerStatus(server.IP)
	}
	if result {
		//log.Println("Ping success:", server.IP)
		dao.UpdateServerStatus(server.ID, 1)
		//GetLinuxConfig(server)
	} else {
		dao.UpdateServerStatus(server.ID, 0)
	}
}

// StartServerStatusCheck 函数用于启动服务器状态检查任务。
//
// 参数:
//
//	interval - 检查间隔时间，单位为秒。
func StartServerStatusCheck(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	var wg sync.WaitGroup
	for {
		select {
		case <-ticker.C:
			serverList := dao.GetServerList()
			for _, server := range serverList {
				wg.Add(1)
				go UpdateServerStatus(server)
				//wg.Wait()
			}
		}
	}
}

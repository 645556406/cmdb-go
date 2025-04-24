package service

import (
	"cmdb-backend/dao"
	"cmdb-backend/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

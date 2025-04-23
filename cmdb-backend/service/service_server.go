package service

import (
	"cmdb-backend/dao"
	"cmdb-backend/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

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

package main

import (
	"cmdb-backend/api"
	"cmdb-backend/service"
	"github.com/gin-contrib/cors"
)

func main() {
	// 初始化引擎
	r := api.NewRouter()
	r.Use(cors.Default())
	r.GET("/ping", service.Ping)
	{
		v1 := r.Group("/api/v1/server")
		v1.GET("/list", service.GetServerList)
		v1.POST("/add", service.AddServer)
		v1.POST("/del", service.DelServer)
		v1.POST("/upd", service.UpdateServer)
		v1.GET("/detail/:id", service.GetServerDetailByID)
		v1.GET("/get/:ip", service.GetServerOneByIP)
	}
	{
		v2 := r.Group("/api/v1/ssh")
		v2.GET("/connect", service.HandleWebSSH)
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

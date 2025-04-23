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
	{
		v1 := r.Group("/api/v1")
		v1.GET("/ping", service.Ping)
		v1.GET("/server/list", service.GetServerList)
		v1.POST("/server/add", service.AddServer)
		v1.POST("/server/del", service.DelServer)
		v1.POST("/server/upd", service.UpdateServer)
		v1.GET("/server/detail/:id", service.GetServerDetailByID)
	}
	//r.GET("/api/v1/server/update", service.UpdateServer)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

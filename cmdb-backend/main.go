package main

import (
	"cmdb-backend/api"
	"cmdb-backend/dao"
	"cmdb-backend/service"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	// 初始化数据库
	dao.NewDB()
	// 启动异步定时任务，每30秒检查一次服务器状态，并更新数据库
	go service.StartServerStatusCheck(10 * time.Second)
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
		v1.GET("/count", service.GetCountServer)
		v1.GET("/count/online", service.GetOnlineCountServer)
		v1.GET("/count/offline", service.GetOfflineCountServer)
	}
	{
		v2 := r.Group("/api/v1/ssh")
		//v2.GET("/connect", service.HandleWebSSH)
		v2.GET("/connect", service.HandleWebSSHSinger)
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

package route

import (
	"cmdb-backend/service"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var R = gin.New()

func InitRouter() {
	R.Use(cors.Default())
	R.GET("/ping", service.Ping)
	{
		v1 := R.Group("/api/v1/server")
		v1.GET("/list", service.GetServerList)
		v1.POST("/add", service.AddServer)
		v1.POST("/del", service.DelServer)
		v1.POST("/upd", service.UpdateServer)
		v1.GET("/detail/:id", service.GetServerDetailByID)
		v1.GET("/get/:ip", service.GetServerOneByIP)
		v1.GET("/count", service.GetCountServer)
		v1.GET("/update", service.HandleWebSocket)
	}
	{
		v2 := R.Group("/api/v1/ssh")
		//v2.GET("/connect", service.HandleWebSSH)
		v2.GET("/connect", service.HandleWebSSHSinger)
	}
	err := R.Run(":8080")
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}

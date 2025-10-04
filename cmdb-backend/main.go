package main

import (
	"cmdb-backend/dao"
	"cmdb-backend/route"
	"log"
)

func main() {
	// 初始化路由
	route.InitRouter()
	// 程序关闭时，关闭数据库
	defer func() {
		if err := dao.CloseDB(); err != nil {
			log.Printf("Warning: error closing database: %v", err)
		}
	}()
}

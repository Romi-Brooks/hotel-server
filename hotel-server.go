package main

import (
	"hotel-server/config"
	"hotel-server/router"
	"log"
)

func main() {
	// 步骤1：先初始化数据库（必须在路由之前执行）
	err := config.InitDB()
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	defer config.DB.Close() // 程序退出时关闭连接池

	// 步骤2：初始化路由并启动服务
	r := router.InitRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

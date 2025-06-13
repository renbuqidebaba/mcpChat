package main

import (
	"backend/config"
	"backend/router"
	"log"
)

func main() {
	r := router.InitRouter()
	config.InitConfig()
	config.InitDatabase()
	// 启动服务器
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
}

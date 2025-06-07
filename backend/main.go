package main

import (
	"backend/router"
	"log"
)

func main() {
	r := router.InitRouter()
	// 启动服务器
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
}

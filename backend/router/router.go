package router

import (
	"backend/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// 配置CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:8081"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))
	api := r.Group("/api")
	{
		api.POST("/chat", service.Chat)
	}
	user := r.Group("/user")
	{
		user.POST("/login", service.UserLogin)
		user.POST("/register", service.UserRegister)
		user.GET("/info", service.UserInfo) // 假设有一个获取用户信息的接口
	}
	return r
}

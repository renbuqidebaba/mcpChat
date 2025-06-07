package service

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Chat(c *gin.Context) {
	var req models.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求格式"})
		return
	}
	msg := models.SentChat(req.Message)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  msg,
	})
}

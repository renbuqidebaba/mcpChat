package test

import (
	"backend/models"
	"testing"
)

func TestSentChat(t *testing.T) {
	models.SentChat("你好，请问你是基于什么模型开发的")
}

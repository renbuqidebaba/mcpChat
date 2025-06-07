package models

import (
	"encoding/json"
	"io/ioutil"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Id     string        `json:"id"`
	Model  string        `json:"model"`
	Choice []ChoicesInfo `json:"choices"`
}

type ChoicesInfo struct {
	Index       int    `json:"index"`
	FinishReson string `json:"finish_reason"`
	Message     Msg    `json:"message"`
}

type Msg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func SentChat(question string) string {
	url := "https://api.siliconflow.cn/v1/chat/completions"
	// Request payload
	payload := map[string]interface{}{
		"model": "Pro/deepseek-ai/DeepSeek-V3-1226",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": question,
			},
		},
		"temperature": 0.7,
		// "tools": []map[string]interface{}{
		// 	{
		// 		"type": "function",
		// 		"function": map[string]interface{}{
		// 			"description": "<string>",
		// 			"name":        "<string>",
		// 			"parameters":  map[string]interface{}{},
		// 			"strict":      false,
		// 		},
		// 	},
		// },
	}
	resp := HttpRequest("POST", url, payload)
	var chatresq ChatResponse
	// Handle response
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &chatresq)
	return chatresq.Choice[0].Message.Content
}

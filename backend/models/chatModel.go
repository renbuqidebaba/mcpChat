package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
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
	url := viper.GetString("llm.url")
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
	}
	resp := HttpRequest("POST", url, payload)
	var chatresq ChatResponse
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &chatresq)
	return chatresq.Choice[0].Message.Content
}

func Getsome() {
	url := "https://api.openai-hub.com/v1/messages"
	method := "POST"

	payload := strings.NewReader(`{` + "\n" +
		`  "model": "claude-3-5-sonnet-20240620",` + "\n" +
		`  "max_tokens": 1024,` + "\n" +
		`  "messages": [` + "\n" +
		`    {` + "\n" +
		`      "role": "user",` + "\n" +
		`      "content": "Hello, world"` + "\n" +
		`    }` + "\n" +
		`  ]` + "\n" +
		`}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer sk-wDvoBfnSp5MRnMvt2MqzWUyQS01II2YmaovleZbQzUbBSDoe")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

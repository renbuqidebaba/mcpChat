package models

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func HttpRequest(method, url string, data interface{}) *http.Response {
	client := &http.Client{}
	var req *http.Request
	if method == "GET" {

	} else {
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Error marshaling data:", err)
			return nil
		}
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Error creating request:", err)
			return nil
		}
	}
	// Set headers
	req.Header.Set("Authorization", "Bearer "+viper.GetString("llm.api_key"))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return nil
	}
	return resp
}

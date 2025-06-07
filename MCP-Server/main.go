package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"MCP 面试助手",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)
	if err := server.ServeStdio(s); err != nil {
		fmt.Println("Error starting MCP server:", err)
	}
}

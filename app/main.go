// Main entry point for the MCP server registering all chromedp tools.
package main

import (
	"context"
	"log"

	"go-browse/app/tools"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "go-browse",
		Version: "v1.0.0",
	}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "launch_browser", Description: "Launch a new browser instance"}, tools.LaunchBrowser)
	mcp.AddTool(server, &mcp.Tool{Name: "navigate", Description: "Navigate to a URL"}, tools.Navigate)
	mcp.AddTool(server, &mcp.Tool{Name: "click", Description: "Click an element"}, tools.Click)
	mcp.AddTool(server, &mcp.Tool{Name: "type_text", Description: "Type text into an element"}, tools.TypeText)
	mcp.AddTool(server, &mcp.Tool{Name: "screenshot", Description: "Take a screenshot"}, tools.Screenshot)
	mcp.AddTool(server, &mcp.Tool{Name: "close_browser", Description: "Close the browser"}, tools.CloseBrowser)

	log.Println("Starting MCP server on stdio")
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

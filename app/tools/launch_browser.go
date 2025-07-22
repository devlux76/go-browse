// Package tools provides MCP tools for browser automation.
package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// LaunchBrowserParams defines parameters for the launch_browser tool.
type LaunchBrowserParams struct{}

// LaunchBrowserResult defines the result for the launch_browser tool.
type LaunchBrowserResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// LaunchBrowser is an MCP tool that initializes the chromedp context and browser.
func LaunchBrowser(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[LaunchBrowserParams]) (*mcp.CallToolResultFor[LaunchBrowserResult], error) {
	InitChromeDPContext()
	if ChromeCtx == nil {
		return &mcp.CallToolResultFor[LaunchBrowserResult]{
			Content: []mcp.Content{},
			StructuredContent: LaunchBrowserResult{
				Success: false,
				Message: "Failed to initialize ChromeDP context",
			},
			IsError: true,
		}, nil
	}
	return &mcp.CallToolResultFor[LaunchBrowserResult]{
		Content: []mcp.Content{},
		StructuredContent: LaunchBrowserResult{
			Success: true,
			Message: "ChromeDP context initialized",
		},
		IsError: false,
	}, nil
}

// Package tools provides MCP tools for browser automation.
package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// StatResult defines the result for the stat tool.
type StatResult struct {
	BrowserConnected bool   `json:"browser_connected"`
	Message          string `json:"message"`
}

// Stat is an MCP tool that checks if the ChromeDP context is initialized and connected.
func Stat(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[struct{}]) (*mcp.CallToolResultFor[StatResult], error) {
	connected := ChromeCtx != nil
	msg := "Browser is not connected"
	if connected {
		msg = "Browser is connected"
	}
	return &mcp.CallToolResultFor[StatResult]{
		Content: []mcp.Content{},
		StructuredContent: StatResult{
			BrowserConnected: connected,
			Message:          msg,
		},
		IsError: false,
	}, nil
}

// Package tools provides MCP tools for browser automation.
package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// CloseBrowserParams defines the input parameters for the close_browser tool.
type CloseBrowserParams struct{}

// CloseBrowserResult defines the output result for the close_browser tool.
type CloseBrowserResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// CloseBrowser closes the current Chrome browser session.
func CloseBrowser(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[CloseBrowserParams]) (*CloseBrowserResult, error) {
	if ChromeCancel != nil {
		ChromeCancel()
		ChromeCtx = nil
		ChromeCancel = nil
		return &CloseBrowserResult{
			Success: true,
			Message: "Browser closed successfully",
		}, nil
	}
	return &CloseBrowserResult{
		Success: false,
		Message: "No browser session to close",
	}, nil
}

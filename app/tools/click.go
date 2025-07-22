// Package tools provides MCP tools for browser automation.
package tools

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ClickParams defines parameters for the click tool.
type ClickParams struct {
	Selector string `json:"selector"`
}

// ClickResult defines the result for the click tool.
type ClickResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Click is an MCP tool that clicks an element specified by selector using the existing chromedp context.
func Click(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ClickParams]) (*mcp.CallToolResultFor[ClickResult], error) {
	if ChromeCtx == nil {
		return &mcp.CallToolResultFor[ClickResult]{
			Content: []mcp.Content{},
			StructuredContent: ClickResult{
				Success: false,
				Message: "ChromeDP context is not initialized",
			},
			IsError: true,
		}, nil
	}
	err := chromedp.Run(ChromeCtx, chromedp.Click(params.Arguments.Selector))
	if err != nil {
		return &mcp.CallToolResultFor[ClickResult]{
			Content: []mcp.Content{},
			StructuredContent: ClickResult{
				Success: false,
				Message: "Click failed: " + err.Error(),
			},
			IsError: true,
		}, nil
	}
	return &mcp.CallToolResultFor[ClickResult]{
		Content: []mcp.Content{},
		StructuredContent: ClickResult{
			Success: true,
			Message: "Click succeeded",
		},
		IsError: false,
	}, nil
}

// Package tools provides MCP tools for browser automation.
package tools

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// NavigateParams defines parameters for the navigate tool.
type NavigateParams struct {
	URL string `json:"url"`
}

// NavigateResult defines the result for the navigate tool.
type NavigateResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Navigate is an MCP tool that navigates the browser to a specified URL using the existing chromedp context.
func Navigate(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[NavigateParams]) (*mcp.CallToolResultFor[NavigateResult], error) {
	if ChromeCtx == nil {
		return &mcp.CallToolResultFor[NavigateResult]{
			Content: []mcp.Content{},
			StructuredContent: NavigateResult{
				Success: false,
				Message: "ChromeDP context is not initialized",
			},
			IsError: true,
		}, nil
	}
	err := chromedp.Run(ChromeCtx, chromedp.Navigate(params.Arguments.URL))
	if err != nil {
		return &mcp.CallToolResultFor[NavigateResult]{
			Content: []mcp.Content{},
			StructuredContent: NavigateResult{
				Success: false,
				Message: "Navigation failed: " + err.Error(),
			},
			IsError: true,
		}, nil
	}
	return &mcp.CallToolResultFor[NavigateResult]{
		Content: []mcp.Content{},
		StructuredContent: NavigateResult{
			Success: true,
			Message: "Navigation succeeded",
		},
		IsError: false,
	}, nil
}

// Package tools provides MCP tools for browser automation.
package tools

import (
	"context"

	"github.com/chromedp/chromedp"
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
	// Actually launch the browser by running a no-op navigation.
	err := chromedp.Run(ChromeCtx, chromedp.Navigate("about:blank"))
	if err != nil {
		return &mcp.CallToolResultFor[LaunchBrowserResult]{
			Content: []mcp.Content{},
			StructuredContent: LaunchBrowserResult{
				Success: false,
				Message: "Failed to launch browser: " + err.Error(),
			},
			IsError: true,
		}, nil
	}
	return &mcp.CallToolResultFor[LaunchBrowserResult]{
		Content: []mcp.Content{},
		StructuredContent: LaunchBrowserResult{
			Success: true,
			Message: "ChromeDP context initialized and browser launched",
		},
		IsError: false,
	}, nil
}

// TestChrome is an MCP tool that tests launching Chrome via chromedp, independent of global context.
// It mimics https://github.com/chromedp/examples/blob/master/remote/main.go
func TestChrome(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[struct{}]) (*mcp.CallToolResultFor[LaunchBrowserResult], error) {
	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, chromedp.DefaultExecAllocatorOptions[:]...)
	defer allocCancel()
	chromeCtx, chromeCancel := chromedp.NewContext(allocCtx)
	defer chromeCancel()
	err := chromedp.Run(chromeCtx, chromedp.Navigate("about:blank"))
	if err != nil {
		return &mcp.CallToolResultFor[LaunchBrowserResult]{
			Content: []mcp.Content{},
			StructuredContent: LaunchBrowserResult{
				Success: false,
				Message: "Failed to launch Chrome: " + err.Error(),
			},
			IsError: true,
		}, nil
	}
	return &mcp.CallToolResultFor[LaunchBrowserResult]{
		Content: []mcp.Content{},
		StructuredContent: LaunchBrowserResult{
			Success: true,
			Message: "Chrome launched and navigated to about:blank",
		},
		IsError: false,
	}, nil
}

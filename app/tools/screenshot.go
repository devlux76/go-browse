// Package tools provides MCP tools for browser automation.
package tools

import (
	"context"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ScreenshotParams defines parameters for the screenshot tool.
type ScreenshotParams struct {
	FilePath string `json:"file_path"`
	Selector string `json:"selector,omitempty"` // Optional: if set, screenshot only this element
}

// ScreenshotResult defines the result for the screenshot tool.
type ScreenshotResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Screenshot is an MCP tool that takes a screenshot using the existing chromedp context.
func Screenshot(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ScreenshotParams]) (*mcp.CallToolResultFor[ScreenshotResult], error) {
	if ChromeCtx == nil {
		return &mcp.CallToolResultFor[ScreenshotResult]{
			StructuredContent: ScreenshotResult{
				Success: false,
				Message: "ChromeDP context is not initialized",
			},
			IsError: true,
		}, nil
	}

	var buf []byte
	var err error
	if params.Arguments.Selector != "" {
		err = chromedp.Run(ChromeCtx,
			chromedp.Screenshot(params.Arguments.Selector, &buf, chromedp.NodeVisible, chromedp.ByQuery),
		)
	} else {
		err = chromedp.Run(ChromeCtx,
			chromedp.FullScreenshot(&buf, 90),
		)
	}
	if err != nil {
		return &mcp.CallToolResultFor[ScreenshotResult]{
			StructuredContent: ScreenshotResult{
				Success: false,
				Message: "Screenshot failed: " + err.Error(),
			},
			IsError: true,
		}, nil
	}

	if writeErr := os.WriteFile(params.Arguments.FilePath, buf, 0644); writeErr != nil {
		return &mcp.CallToolResultFor[ScreenshotResult]{
			StructuredContent: ScreenshotResult{
				Success: false,
				Message: "Failed to save screenshot: " + writeErr.Error(),
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResultFor[ScreenshotResult]{
		StructuredContent: ScreenshotResult{
			Success: true,
			Message: "Screenshot saved to " + params.Arguments.FilePath,
		},
		IsError: false,
	}, nil
}

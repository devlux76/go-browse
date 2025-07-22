// Package tools provides MCP tools for browser automation.
package tools

import (
	"context"
	"encoding/json"
	"log"
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
	Content []byte `json:"content,omitempty"` // Optional: screenshot content if needed
}

// Screenshot is an MCP tool that takes a screenshot using the existing chromedp context.
func Screenshot(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ScreenshotParams]) (*mcp.CallToolResultFor[ScreenshotResult], error) {
	if ChromeCtx == nil {
		resp := &mcp.CallToolResultFor[ScreenshotResult]{
			Content: []mcp.Content{},
			StructuredContent: ScreenshotResult{
				Success: false,
				Message: "ChromeDP context is not initialized",
			},
			IsError: true,
		}
		if b, err := json.Marshal(resp); err == nil {
			log.Printf("Screenshot tool response: %s", string(b))
		}
		return resp, nil
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
		resp := &mcp.CallToolResultFor[ScreenshotResult]{
			Content: []mcp.Content{},
			StructuredContent: ScreenshotResult{
				Success: false,
				Message: "Screenshot failed: " + err.Error(),
			},
			IsError: true,
		}
		if b, err := json.Marshal(resp); err == nil {
			log.Printf("Screenshot tool response: %s", string(b))
		}
		return resp, nil
	}

	if writeErr := os.WriteFile(params.Arguments.FilePath, buf, 0644); writeErr != nil {
		resp := &mcp.CallToolResultFor[ScreenshotResult]{
			Content: []mcp.Content{},
			StructuredContent: ScreenshotResult{
				Success: false,
				Message: "Failed to save screenshot: " + writeErr.Error(),
			},
			IsError: true,
		}
		if b, err := json.Marshal(resp); err == nil {
			log.Printf("Screenshot tool response: %s", string(b))
		}
		return resp, nil
	}

	resp := &mcp.CallToolResultFor[ScreenshotResult]{
		Content: []mcp.Content{},
		StructuredContent: ScreenshotResult{
			Success: true,
			Message: "Screenshot saved to " + params.Arguments.FilePath,
		},
		IsError: false,
	}
	if b, err := json.Marshal(resp); err == nil {
		log.Printf("Screenshot tool response: %s", string(b))
	}
	return resp, nil
}

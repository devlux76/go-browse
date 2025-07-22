// Package tools provides MCP tools for browser automation.
package tools

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// TypeTextParams defines parameters for the type_text tool.
type TypeTextParams struct {
	Selector string `json:"selector"`
	Text     string `json:"text"`
}

// TypeTextResult defines the result for the type_text tool.
type TypeTextResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// TypeText is an MCP tool that types text into an element specified by selector using the existing chromedp context.
func TypeText(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[TypeTextParams]) (*mcp.CallToolResultFor[TypeTextResult], error) {
	if ChromeCtx == nil {
		return &mcp.CallToolResultFor[TypeTextResult]{
			StructuredContent: TypeTextResult{
				Success: false,
				Message: "ChromeDP context is not initialized",
			},
			IsError: true,
		}, nil
	}
	err := chromedp.Run(ChromeCtx, chromedp.SendKeys(params.Arguments.Selector, params.Arguments.Text))
	if err != nil {
		return &mcp.CallToolResultFor[TypeTextResult]{
			StructuredContent: TypeTextResult{
				Success: false,
				Message: "TypeText failed: " + err.Error(),
			},
			IsError: true,
		}, nil
	}
	return &mcp.CallToolResultFor[TypeTextResult]{
		StructuredContent: TypeTextResult{
			Success: true,
			Message: "TypeText succeeded",
		},
		IsError: false,
	}, nil
}

package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// chromedp helpers for browser automation

// Launches a new Chrome instance and returns a context
func launchBrowser() (context.Context, context.CancelFunc, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	ctx, _ = context.WithTimeout(ctx, 60*time.Second)
	return ctx, cancel, nil
}

// Navigates to a URL
func navigate(ctx context.Context, url string) error {
	return chromedp.Run(ctx, chromedp.Navigate(url))
}

// Clicks an element by selector
func click(ctx context.Context, selector string) error {
	return chromedp.Run(ctx, chromedp.Click(selector))
}

// Types text into an element by selector
func typeText(ctx context.Context, selector, text string) error {
	return chromedp.Run(ctx, chromedp.SendKeys(selector, text))
}

// Parameters for chrome_devtools tool
type ChromeDevToolsParams struct {
	Action   string `json:"action" jsonschema:"Action to perform: navigate, click, type, screenshot"`
	URL      string `json:"url,omitempty" jsonschema:"URL to navigate to"`
	Selector string `json:"selector,omitempty" jsonschema:"CSS selector for click/type actions"`
	Text     string `json:"text,omitempty" jsonschema:"Text to type (for type action)"`
}

// Takes a screenshot of the current page
func screenshot(ctx context.Context) ([]byte, error) {
	var buf []byte
	err := chromedp.Run(ctx, chromedp.FullScreenshot(&buf, 90))
	return buf, err
}

// Gets browser console logs (basic example)
func getConsoleLogs(ctx context.Context) ([]string, error) {
	// chromedp does not provide direct log capture, would need to use ListenTarget
	// This is a placeholder for future extension
	return []string{}, nil
}

type HiParams struct {
	Name string `json:"name" jsonschema:"the name of the person to greet"`
}

func SayHi(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[HiParams]) (*mcp.CallToolResultFor[any], error) {
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: "Hi " + params.Arguments.Name}},
	}, nil
}

func main() {
	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "say hi"}, SayHi)
	// Run the server over stdin/stdout, until the client disconnects
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}

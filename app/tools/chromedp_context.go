// Package tools provides global chromedp context and cancel function.
package tools

import (
	"context"

	"github.com/chromedp/chromedp"
)

var (
	// ChromeCtx is the global chromedp context.
	ChromeCtx context.Context
	// ChromeCancel is the global cancel function for chromedp context.
	ChromeCancel context.CancelFunc
)

func InitChromeDPContext() {
	ChromeCtx, ChromeCancel = chromedp.NewContext(context.Background())
}

// Command emulate is a chromedp example demonstrating how to emulate a
// specific device.
package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
			chromedp.Flag("enable-automation", false),
			chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"),
		)...)
	defer cancel()

	ctx, _ = chromedp.NewContext(ctx)

	// hasLoginSel := `//section[@id="contents"]`
	// emailSel := `//*[@id="logo"]/img`
	// hasLogin := false

	waitIdx := 0
	urlstr := ""
	if err := chromedp.Run(
		ctx,
		chromedp.Navigate(`about:blank`),
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(map[string]interface{}{
			"cache-control": "no-store, no-cache, must-revalidate, post-check=0, pre-check=0",
		})),
		chromedp.Navigate(`https://www.baidu.com`, chromedp.NavigateNoWait),
		chromedp.WaitOneOf(&waitIdx,
			// chromedp.WaitVisible(hasLoginSel),
			// chromedp.WaitVisible(emailSel),
			chromedp.WaitNavigate(chromedp.NavigateWaitEventLoadEventFired),
			chromedp.WaitLocation(&urlstr),
		),
		chromedp.Sleep(300*time.Second),
	); err != nil {
		log.Fatal(err)
	}

	log.Println("ret", waitIdx)
	log.Println("urlstr", urlstr)
}

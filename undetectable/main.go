// Command emulate is a chromedp example demonstrating how to emulate a
// specific device.
package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true),
			chromedp.Flag("enable-automation", false),
			chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"),
		)...)
	defer cancel()

	ctx, _ = chromedp.NewContext(ctx)

	var b1, b2, b3 []byte
	if err := chromedp.Run(ctx,
		chromedp.Undetectable(chromedp.BypassIframeTest(true)),
		chromedp.Navigate(`https://intoli.com/blog/making-chrome-headless-undetectable/chrome-headless-test.html`),
		chromedp.Sleep(1*time.Second),
		chromedp.CaptureScreenshot(&b1),
		chromedp.Navigate(`https://intoli.com/blog/not-possible-to-block-chrome-headless/chrome-headless-test.html`),
		chromedp.Sleep(1*time.Second),
		chromedp.CaptureScreenshot(&b2),
		chromedp.EmulateViewport(1920, 3000),
		chromedp.Navigate(`https://antoinevastel.com/bots/`),
		chromedp.Sleep(2*time.Second),
		chromedp.CaptureScreenshot(&b3),
	); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("screenshot1.png", b1, 0644); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("screenshot2.png", b2, 0644); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("screenshot3.png", b3, 0644); err != nil {
		log.Fatal(err)
	}
}

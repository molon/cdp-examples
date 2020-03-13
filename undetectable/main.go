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
	ctx := context.Background()
	allocOpts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("window-position", "0,0"),
		// chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"),
		chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36"),
		// chromedp.ProxyServer("http://127.0.0.1:8888"),
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, allocOpts...)
	defer allocCancel()

	bctx, _ := chromedp.NewContext(allocCtx) //, chromedp.WithDebugf(log.Printf))

	var b1, b2, b3, b4, b5, b6 []byte
	if err := chromedp.Run(
		bctx,
		// chromedp.Undetectable(chromedp.BypassIframeTest(false)), // 这个只能作用于当前页而已
		chromedp.Undetectable(chromedp.BypassIframeTest(true)), // 这个只能作用于当前页而已
		chromedp.Navigate(`https://www.footaction.com/`),
		chromedp.Sleep(2000*time.Second),

		// chromedp.EmulateViewport(1920, 9000),
		// chromedp.Navigate(`https://bot.sannysoft.com/`),
		// chromedp.Sleep(2*time.Second),
		// chromedp.CaptureScreenshot(&b1),

		// chromedp.Navigate(`https://recaptcha-demo.appspot.com/recaptcha-v3-request-scores.php`),
		// chromedp.Sleep(5000*time.Second),
		// chromedp.CaptureScreenshot(&b1),

		// chromedp.Navigate(`https://intoli.com/blog/making-chrome-headless-undetectable/chrome-headless-test.html`),
		// chromedp.Sleep(2*time.Second),
		// chromedp.CaptureScreenshot(&b2),

		// chromedp.Navigate(`https://intoli.com/blog/not-possible-to-block-chrome-headless/chrome-headless-test.html`),
		// chromedp.Sleep(2*time.Second),
		// chromedp.CaptureScreenshot(&b3),

		// chromedp.Navigate(`https://antoinevastel.com/bots/`),
		// chromedp.Sleep(2*time.Second),
		// chromedp.CaptureScreenshot(&b4),

		// chromedp.Navigate(`https://infosimples.github.io/detect-headless/`), // 这个也没过去，需要点击OK
		// chromedp.Sleep(2*time.Second),
		// chromedp.CaptureScreenshot(&b5),

		// chromedp.Navigate(`https://arh.antoinevastel.com/bots/areyouheadless`), // 这个没过去 资料：https://www.tenantbase.com/tech/blog/cat-and-mouse/
		// chromedp.Sleep(2*time.Second),
		// chromedp.CaptureScreenshot(&b6),
	); err != nil {
		log.Fatal(err)
	}

	if len(b1) > 0 {
		if err := ioutil.WriteFile("screenshot1.png", b1, 0644); err != nil {
			log.Fatal(err)
		}
	}
	if len(b2) > 0 {
		if err := ioutil.WriteFile("screenshot2.png", b2, 0644); err != nil {
			log.Fatal(err)
		}
	}
	if len(b3) > 0 {
		if err := ioutil.WriteFile("screenshot3.png", b3, 0644); err != nil {
			log.Fatal(err)
		}
	}
	if len(b4) > 0 {
		if err := ioutil.WriteFile("screenshot4.png", b4, 0644); err != nil {
			log.Fatal(err)
		}
	}
	if len(b5) > 0 {
		if err := ioutil.WriteFile("screenshot5.png", b5, 0644); err != nil {
			log.Fatal(err)
		}
	}
	if len(b6) > 0 {
		if err := ioutil.WriteFile("screenshot6.png", b6, 0644); err != nil {
			log.Fatal(err)
		}
	}
}

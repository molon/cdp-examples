// Command emulate is a chromedp example demonstrating how to emulate a
// specific device.
package main

import (
	"context"
	"log"

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
		chromedp.Navigate(`https://www.solebox.com/index.php?actcontrol=payment&cl=payment&fnc=validatepayment&paymentid=globalpaypal&lang=0&stoken=&userfrom=`, chromedp.NavigateNoWait),
		chromedp.WaitOneOf(&waitIdx,
			// chromedp.WaitVisible(hasLoginSel),
			// chromedp.WaitVisible(emailSel),
			chromedp.WaitNavigate(chromedp.NavigateWaitEventLoadEventFired),
			chromedp.WaitLocation(&urlstr),
		),
	); err != nil {
		log.Fatal(err)
	}

	log.Println("ret", waitIdx)
	log.Println("urlstr", urlstr)
}

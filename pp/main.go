package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
			chromedp.Flag("enable-automation", false),
		)...)
	defer cancel()

	ctx, _ = chromedp.NewContext(ctx, chromedp.WithDebugf(func(f string, v ...interface{}) {
		log.Printf(f, v...)
	}))
	username := `xxx`
	password := `yyy`

	emailSel := `//input[@id="email"]`
	passSel := `//input[@id="password"]`
	nextSel := `//button[@id='btnNext']`
	loginSel := `//button[@id='btnLogin']`
	hasLoginSel := `//section[@id="contents"]`
	if err := chromedp.Run(
		ctx,
		chromedp.Navigate(`https://www.paypal.com/myaccount/summary/`),
		chromedp.SetValue(emailSel, username, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),
		chromedp.Click(nextSel, chromedp.NodeVisible),
		chromedp.WaitNotVisible(nextSel),
		chromedp.Sleep(1*time.Second),
		chromedp.SendKeys(passSel, password, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),
		chromedp.Click(loginSel, chromedp.NodeVisible),
		chromedp.WaitVisible(hasLoginSel),
	); err != nil {
		log.Fatal(err)
	}
}

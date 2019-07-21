package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
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

	chromedp.ListenTarget(ctx, func(event interface{}) {
		switch ev := event.(type) {
		case *network.EventRequestWillBeSent:
			go func() {
				log.Println(ev.RequestID)

				c := chromedp.FromContext(ctx)
				body, err := network.GetResponseBody(ev.RequestID).Do(cdp.WithExecutor(ctx, c.Target))
				if err != nil {
					log.Println("getting body error: ", err)
					return
				}
				if err = ioutil.WriteFile(ev.RequestID.String(), body, 0644); err != nil {
					log.Fatal(err)
				}
			}()
		}
	})
	err := chromedp.Run(ctx, tasks())
	if err != nil {
		log.Fatal(err)
	}
}

func tasks() chromedp.Tasks {
	return chromedp.Tasks{
		network.Enable(),
		chromedp.Navigate("https://www.baidu.com/"),
		chromedp.Sleep(time.Second * 100),
	}
}

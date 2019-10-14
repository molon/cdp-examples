package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/molon/pkg/util"
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
		case *fetch.EventRequestPaused:
			go func() {
				c := chromedp.FromContext(ctx)
				ctx := cdp.WithExecutor(ctx, c.Target)

				if ev.Request.URL != "http://home.baidu.com/" {
					fetch.ContinueRequest(ev.RequestID).Do(ctx)
					return
				}

				util.PrintJson(ev)

				// 删除俩cookie试试
				err := network.DeleteCookies("BDORZ").WithDomain(".baidu.com").Do(ctx)
				if err != nil {
					log.Printf("1.err:%+v\n", err)
				}

				err = network.DeleteCookies("BAIDUID").WithDomain(".baidu.com").Do(ctx)
				if err != nil {
					log.Printf("3.err:%+v\n", err)
				}

				cookies, err := network.GetCookies().WithUrls([]string{ev.Request.URL}).Do(ctx)
				if err != nil {
					log.Printf("2.err:%+v\n", err)
				}
				util.PrintJson(cookies)

				fetch.ContinueRequest(ev.RequestID).Do(ctx)

				// headers := []*fetch.HeaderEntry{}
				// for k, v := range ev.Request.Headers {
				// 	value := fmt.Sprint(v)
				// 	if k == "Cookie" {
				// 		value = "BAIDUID=01C8BB7BA07A34837B01B8A243832C93:FG=1;"
				// 	}
				// 	headers = append(headers, &fetch.HeaderEntry{
				// 		Name:  k,
				// 		Value: value,
				// 	})
				// }
				// headers = append(headers, &fetch.HeaderEntry{
				// 	Name:  "fake",
				// 	Value: "fake_value",
				// })

				// if ev.Request.URL == "http://home.baidu.com/" {
				// 	util.PrintJson(headers)
				// }

				// fetch.ContinueRequest(ev.RequestID).Do(ctx)
			}()
			// go func() {
			// 	log.Println(ev.RequestID)

			// 	c := chromedp.FromContext(ctx)
			// 	body, err := network.GetResponseBody(ev.RequestID).Do(cdp.WithExecutor(ctx, c.Target))
			// 	if err != nil {
			// 		log.Println("getting body error: ", err)
			// 		return
			// 	}
			// 	if err = ioutil.WriteFile(ev.RequestID.String(), body, 0644); err != nil {
			// 		log.Fatal(err)
			// 	}
			// }()
		}
	})
	err := chromedp.Run(ctx, tasks())
	if err != nil {
		log.Fatal(err)
	}
}

func tasks() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Undetectable(),
		fetch.Enable(),
		// network.Enable(),
		chromedp.Navigate("https://www.baidu.com/"),
		chromedp.Sleep(time.Second * 100),
	}
}

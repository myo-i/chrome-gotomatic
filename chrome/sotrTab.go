package chrome

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func TabSort() {
	// // create chrome instance
	// ctx, cancel := chromedp.NewContext(
	// 	context.Background(),
	// 	// chromedp.WithDebugf(log.Printf),
	// )
	// defer cancel()

	// // create a timeout
	// ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	// defer cancel()

	// // navigate to a page, wait for an element, click
	// var example string
	// err := chromedp.Run(ctx,
	// 	chromedp.Navigate(`https://pkg.go.dev/time`),
	// 	// wait for footer element is visible (ie, page is loaded)
	// 	chromedp.WaitVisible(`body > footer`),
	// 	// find and click "Example" link
	// 	chromedp.Click(`#example-After`, chromedp.NodeVisible),
	// 	// retrieve the text of the textarea
	// 	chromedp.Value(`#example-After textarea`, &example),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Go's time.After example:\n%s", example)

	// opts := append(chromedp.DefaultExecAllocatorOptions[:],
	// 	chromedp.DisableGPU,
	// 	chromedp.UserDataDir("C:\\Users\\PC_User\\AppData\\Local\\Google\\Chrome\\User Data"),
	// 	chromedp.Flag("headless", false),
	// 	chromedp.Flag("enable-automation", false),
	// 	chromedp.Flag("restore-on-startup", false),
	// )
	// allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	// ctx, _ := chromedp.NewContext(allocCtx)
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	if err := chromedp.Run(ctx, chromedp.Navigate("https://finance.yahoo.co.jp/")); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }

	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithErrorf(log.Printf))
	defer cancel()
	var nodes []*cdp.Node
	selector := "#main ul li a"
	pageURL := "https://notepad-plus-plus.org/downloads/"
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(pageURL),
		chromedp.WaitReady(selector),
		chromedp.Nodes(selector, &nodes),
	}); err != nil {
		log.Fatalf("Failed1: ", err)
	}
	f := func(ctx context.Context, url string) {
		clone, cancel := chromedp.NewContext(ctx)
		defer cancel()
		fmt.Printf("%s is opening in a new tab\n", url)

		if err := chromedp.Run(clone, chromedp.Navigate(url)); err != nil {
			// do something nice with you errors!
			log.Fatalf("Failed2: ", err)
		}
	}
	for _, n := range nodes {
		u := n.AttributeValue("href")
		go f(ctx, u)
	}

	// for i := 0; i < 5; i++ {
	// 	var res *runtime.RemoteObject
	// 	if err := chromedp.Run(ctx, chromedp.Evaluate(`window.open("about:blank", "", "resizable,scrollbars,status")`, &res)); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	targets, err := chromedp.Targets(ctx)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	for _, t := range targets {
	// 		if !t.Attached {
	// 			newCtx, _ := chromedp.NewContext(ctx, chromedp.WithTargetID(t.TargetID))
	// 			if err := chromedp.Run(newCtx, chromedp.Navigate("https://example.com")); err != nil {
	// 				log.Fatalln(err)
	// 			}
	// 		}
	// 	}
	// }

	time.Sleep(time.Minute)
}

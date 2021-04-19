package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

//我创建了一个context.Context并且为它添加了超时设置，当超时时间到了，”ctx”将接收到channel的超时警告。
//WithTimeout同样也会返回一个取消的方法，但是我们不需要，所以用 “_” 来忽略掉了。
//下面的search() 方法的有context,search path,和 search pattern。最后把找到的文件和数量输出。

func main1() {
	//f1sd()
	f3()
	fmt.Println("waitGroup完成")
}

// waitGroup
func f1sd() {

	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

func f3() {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else if err != nil {
		fmt.Println("failed fetched all URLs.", err)
	}

}
func f4() {
	//zap.L().Panic("group_test: doomed", zap.Error(err))
	errDoom := errors.New("group_test: doomed")

	cases := []struct {
		errs []error
		want error
	}{
		{want: nil},
		{errs: []error{nil}, want: nil},
		{errs: []error{errDoom}, want: errDoom},
		{errs: []error{nil, errDoom}, want: errDoom},
	}

	for _, tc := range cases {
		g, ctx := errgroup.WithContext(context.Background())

		for _, err := range tc.errs {
			err := err
			g.Go(func() error {
				zap.L().Error("group_test: doomed", zap.Error(err))
				//log.Error(err) // 当此时的err = nil 时，g.Go不会将 为nil 的 err 放入g.err中
				return err
			})
		}
		err := g.Wait() // 这里等待所有Go跑完即add==0时，此处的err是g.err的信息。
		//log.Error(err)
		zap.L().Error("g.Wait", zap.Error(err))
		//log.Error(tc.want)
		zap.L().Error("want", zap.Error(tc.want))
		if err != tc.want {
			//t.Errorf("after %T.Go(func() error { return err }) for err in %v\n"+
			//	"g.Wait() = %v; want %v",
			//	g, tc.errs, err, tc.want)
			fmt.Printf("after %T.Go(func() error { return err }) for err in %v\n"+
				"g.Wait() = %v; want %v",
				g, tc.errs, err, tc.want)
		}

		canceled := false
		select {
		case <-ctx.Done():
			// 由于上文中内部调用了cancel(),所以会有Done()接受到了消息
			// returns an error or ctx.Done is closed
			// 在当前工作完成或者上下文被取消之后关闭
			canceled = true
		default:
		}
		if !canceled {
			//t.Errorf("after %T.Go(func() error { return err }) for err in %v\n"+
			//	"ctx.Done() was not closed",
			//	g, tc.errs)
			fmt.Printf("after %T.Go(func() error { return err }) for err in %v\n"+
				"ctx.Done() was not closed",
				g, tc.errs)
		}
	}
}

func f2() {

	//var g errgroup.Group
	//var urls = []string{
	//	"http://www.golang.org/",
	//	"http://www.google.com/",
	//	"http://www.somestupidname.com/",
	//}
	//for _, url := range urls {
	//	// Launch a goroutine to fetch the URL.
	//	url := url // https://golang.org/doc/faq#closures_and_goroutines
	//	g.Go(func(ctx context.Context) error {
	//		// Fetch the URL.
	//		resp, err := http.Get(url)
	//		if err == nil {
	//			resp.Body.Close()
	//		}
	//		return err
	//	})
	//}
	//// Wait for all HTTP fetches to complete.
	//if err := g.Wait(); err == nil {
	//	fmt.Println("Successfully fetched all URLs.")
	//} else {
	//	fmt.Println(" fetched  URLs. failed")
	//}

}

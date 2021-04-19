package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/neilotoole/errgroup"
)

func fn() {

	g, _ := errgroup.WithContextN(context.Background(), 1, 1)

	var count int64
	g.Go(func() error {
		time.Sleep(time.Second) //睡眠5秒，把这个goroutine占住
		return nil
	})

	total := 1000

	for i := 0; i < total; i++ { // 并发一万个goroutine执行子任务，理论上这些子任务都会加入到Group的待处理列表中
		go func() {
			g.Go(func() error {
				atomic.AddInt64(&count, 1)
				return nil
			})
		}()
	}

	// 等待所有的子任务完成。理论上10001个子任务都会被完成
	if err := g.Wait(); err != nil {
		panic(err)
	}

	got := atomic.LoadInt64(&count)
	if got != int64(total) {
		panic(fmt.Sprintf("expect %d but got %d", total, got))
	}

}

func main() {
	fn()
}

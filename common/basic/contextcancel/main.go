package main

import (
	"context"
	"fmt"
	"time"
)

func Timeoutfunc() { //这个 { 不能另起一行
	ctx, _ := context.WithTimeout(context.TODO(), 2*time.Second)
	go func(ctx context.Context) {
		fmt.Printf("begin func1\n")
		go func(ctx context.Context) {
			ch := make(chan error, 0)
			go func(ctx context.Context) {
				fmt.Printf("begin func2\n")
				time.Sleep(5 * time.Second)
				if ctx.Err() != nil {
					dealLine, dead := ctx.Deadline()
					fmt.Printf("do not call end func2,deadLine:%v,dead:%v\n", dealLine, dead)
					return
				}
				fmt.Printf("end func2\n")
				ch <- nil
			}(ctx)
			select {
			case <-ch:
				fmt.Printf("ok")
			case <-ctx.Done():
				fmt.Printf("err:%v\n", ctx.Err())
			}
		}(ctx)
		fmt.Printf("end func1\n")
	}(ctx)
	time.Sleep(10 * time.Second)
}

func main() {
	Timeoutfunc()
	//ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	//defer cancel()
	//go task(ctx)
	//time.Sleep(time.Second * 10)
}

func task(ctx context.Context) {
	ch := make(chan struct{}, 0)
	go func() {
		// 模拟4秒耗时任务
		time.Sleep(time.Second * 4)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}

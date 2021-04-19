package main

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/facebookgo/errgroup"
)

func fn() {

	var count int64

	var g errgroup.Group
	g.Add(1001)

	// 启动第一个子任务,它执行成功
	go func() {
		time.Sleep(2 * time.Second) //睡眠5秒，把这个goroutine占住
		fmt.Println("exec #1")
		g.Done()
	}()

	total := 1000

	for i := 0; i < total; i++ { // 并发一万个goroutine执行子任务，理论上这些子任务都会加入到Group的待处理列表中
		go func() {
			atomic.AddInt64(&count, 1)

			g.Done()
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

	fmt.Sprintf("expect %d and got %d", total, got)

}
func main() {
	fn()
}

func main1() {
	var g errgroup.Group
	g.Add(4)

	// 启动第一个子任务,它执行成功
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		g.Done()
	}()

	// 启动第二个子任务，它执行失败
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		g.Error(errors.New("failed to exec #2"))
		g.Done()
	}()

	// 启动第三个子任务，它执行成功
	go func() {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		g.Done()
	}()

	// 启动第四个子任务，它执行失败
	go func() {
		time.Sleep(7 * time.Second)
		fmt.Println("exec #4")
		g.Error(errors.New("failed to exec #4"))
		g.Done()
	}()

	// 等待所有的goroutine完成，并检查error
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Println("failed:", err)
	}

}

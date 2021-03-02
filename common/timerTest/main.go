package main

import (
	"fmt"

	"time"
)

func main() {
	TimeAfter()
	//t := time.NewTimer(time.Second * 2)
	//defer t.Stop()
	//for {
	//	<-t.C
	//	fmt.Println("timer running...")
	//	// 需要重置Reset 使 t 重新开始计时
	//	t.Reset(time.Second * 2)
	//}
}
func timeTest() {
	t := time.NewTimer(time.Second * 2)

	ch := make(chan bool)
	go func(t *time.Timer) {
		defer t.Stop()
		for {
			select {
			case <-t.C:
				fmt.Println("timer running....")
				// 需要重置Reset 使 t 重新开始计时
				t.Reset(time.Second * 2)
			case stop := <-ch:
				if stop {
					fmt.Println("timer Stop")
					return
				}
			}
		}
	}(t)
	time.Sleep(10 * time.Second)
	ch <- true
	close(ch)
	time.Sleep(1 * time.Second)
}
func tickerTest() {
	t := time.NewTicker(time.Second * 2)
	defer t.Stop()
	for {
		<-t.C
		fmt.Println("Ticker running...")
	}
}
func tickerTest2() {

	ticker := time.NewTicker(2 * time.Second)
	ch := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker running...")
			case stop := <-ch:
				if stop {
					fmt.Println("Ticker Stop")
					return
				}
			}
		}
	}(ticker)
	time.Sleep(10 * time.Second)
	ch <- true
	close(ch)
	time.Sleep(1 * time.Second)
}

func TimeAfter() {
	t := time.After(time.Second * 3)
	fmt.Printf("t type=%T\n", t)
	//阻塞3秒
	fmt.Println("t=", <-t)
}

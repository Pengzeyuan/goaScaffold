package main

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

//func main() {
//	a := runtime.hchan{}
//
//}

func main() {
	var closing = make(chan struct{})
	var closed = make(chan struct{})

	go func() {
		// 模拟业务处理
		for {
			select {
			case <-closing:
				return
			default:
				// ....... 业务计算
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	// 执行退出之前的清理动作
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("清理超时，不等了")
	}
	fmt.Println("优雅退出")
}

func doCleanup(closed chan struct{}) {
	time.Sleep((time.Minute))
	close(closed)
}

type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch         // 取得令牌    从自己的chan中获得令牌
		fmt.Println((id + 1)) // id从1开始
		time.Sleep(time.Second)
		nextCh <- token // 传递
	}
}

//func main() {
//	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
//
//	// 创建4个worker
//	for i := 0; i < 4; i++ {
//		go newWorker(i, chs[i], chs[(i+1)%4]) //  从0号取  传给  1号
//	}
//
//	//首先把令牌交给第一个worker
//	chs[0] <- struct{}{}
//
//	select {}
//}

//func main() {
//	var ch1 = make(chan int, 10)
//	var ch2 = make(chan int, 10)
//
//	ch1 <- 100
//	ch1 <- 300
//	ch2 <- 200
//	// 创建SelectCase
//	var cases = createCases(ch1, ch2)
//
//	// 执行10次select
//	for i := 0; i < 10; i++ {
//		chosen, recv, ok := reflect.Select(cases)
//		if recv.IsValid() { // recv case
//			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
//		} else { // send case
//			fmt.Println("send:", cases[chosen].Dir, ok)
//		}
//	}
//}

func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	// 创建send case
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}

	return cases
}

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

var (
	a = c + b // == 9
	b = f()   // == 4
	c = f()   // == 5
	d = 3     // == 5 全部初始化完成后
)

func main() {
	var a, b int32 = 0, 0

	go func() {
		atomic.StoreInt32(&a, 1)
		atomic.StoreInt32(&b, 1)
	}()

	for atomic.LoadInt32(&b) == 0 {
		runtime.Gosched()

	}

	fmt.Println(atomic.LoadInt32(&a))
}
func f() int {
	d++
	return d
}

type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

func main1() {
	go setup()
	for g == nil {
	}
	print(g.msg)
}

package main

import "fmt"

func main() {
	Contextfn()
	//funcName()
}

func funcName() {
	var x interface{}
	var y = 10
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是：%T", i)
	case int:
		fmt.Printf("x是 int 类型")
	case float64:
		fmt.Printf("x是 float64 类型")
	case func(int) float64:
		fmt.Printf("x是func(int)类型")
	case bool, string:
		fmt.Printf("x是bool或者string类型")
	default:
		fmt.Printf("未知型")
	}
}

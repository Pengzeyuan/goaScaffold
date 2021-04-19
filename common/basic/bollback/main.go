package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/vardius/gollback"
)

func main() {
	testrace()
	//mainretry()
	//funcName()
}

func mainretry() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 尝试5次，或者超时返回
	res, err := gollback.Retry(ctx, 5, func(ctx context.Context) (interface{}, error) {
		return nil, errors.New("failed")
	})

	fmt.Println(res) // 输出结果
	fmt.Println(err) // 输出错误信息
}

func testrace() {
	rs, errs := gollback.Race( // 调用All方法
		context.Background(),
		func(ctx context.Context) (interface{}, error) {
			time.Sleep(3 * time.Second)
			return 1, nil // 第一个任务没有错误，返回1
		},
		func(ctx context.Context) (interface{}, error) {
			return 22, errors.New("failed#2") // 第二个任务返回一个错误
		},
		func(ctx context.Context) (interface{}, error) {
			return 3, nil // 第三个任务没有错误，返回3
		},
		func(ctx context.Context) (interface{}, error) {
			return 44, errors.New("failed#4") // 第四个任务返回一个错误
		},
	)

	fmt.Println(rs)   // 输出子任务的结果
	fmt.Println(errs) // 输出子任务的错误信息
}

func funcName() {
	rs, errs := gollback.All( // 调用All方法
		context.Background(),
		func(ctx context.Context) (interface{}, error) {
			time.Sleep(3 * time.Second)
			return 1, nil // 第一个任务没有错误，返回1
		},
		func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("failed#2") // 第二个任务返回一个错误
		},
		func(ctx context.Context) (interface{}, error) {
			return 3, nil // 第三个任务没有错误，返回3
		},
		func(ctx context.Context) (interface{}, error) {
			return nil, errors.New("failed#4") // 第四个任务返回一个错误
		},
	)

	fmt.Println(rs)   // 输出子任务的结果
	fmt.Println(errs) // 输出子任务的错误信息
}

package main

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sync/errgroup"
)

func mainerrgroupSimpleTest() {
	var g errgroup.Group

	// 启动第一个子任务,它执行成功
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		return nil
	})
	// 启动第二个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		return errors.New("failed to exec #2")
	})

	// 启动第三个子任务，它执行成功
	g.Go(func() error {
		time.Sleep(3 * time.Second)
		fmt.Println("exec #3")
		return nil
	})
	// 等待三个任务都完成
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Println("failed:", err)
	}
}

func mainerrgroupRecordList() { //errgroupRecordList
	var g errgroup.Group
	var result = make([]error, 4)

	// 启动第一个子任务,它执行成功
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		result[0] = nil // 保存成功或者失败的结果
		return nil
	})

	// 启动第二个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")

		result[1] = errors.New("failed to exec #2") // 保存成功或者失败的结果
		return result[1]
	})

	// 启动第三个子任务，它执行成功
	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		result[2] = nil // 保存成功或者失败的结果
		return nil
	})

	// 启动第四个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #4")

		result[3] = errors.New("failed to exec #4") // 保存成功或者失败的结果
		return result[3]
	})

	if err := g.Wait(); err == nil {
		fmt.Printf("Successfully exec all. result: %v\n", result)
	} else {
		fmt.Printf("failed: %v\n", result)
	}
}

// 一个多阶段的pipeline.使用有限的goroutine计算每个文件的md5值.

// Pipeline demonstrates the use of a Group to implement a multi-stage
// pipeline: a version of the MD5All function with bounded parallelism from
// https://blog.golang.org/pipelines.
func main() {
	m, err := MD5All(context.Background(), ".")
	if err != nil {
		log.Fatal(err)
	}

	for k, sum := range m {
		fmt.Printf("%s:\t%x\n", k, sum)
	}
}

type result struct {
	path string
	sum  [md5.Size]byte
}

// 遍历根目录下所有的文件和子文件夹,计算它们的md5的值.

// MD5All reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents. If the directory walk
// fails or any read operation fails, MD5All returns an error.
func MD5All(ctx context.Context, root string) (map[string][md5.Size]byte, error) {
	// ctx is canceled when g.Wait() returns. When this version of MD5All returns
	// - even in case of error! - we know that all of the goroutines have finished
	// and the memory they were using can be garbage-collected.
	g, ctx := errgroup.WithContext(ctx)
	paths := make(chan string) // 文件路径channel

	g.Go(func() error {
		defer close(paths) // 遍历完关闭paths chan
		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			//...... //将文件路径放入到paths
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	})

	// 启动20个goroutine执行计算md5的任务，计算的文件由上一阶段的文件遍历子任务生成.
	// Start a fixed number of goroutines to read and digest files.
	c := make(chan result)
	const numDigesters = 20
	for i := 0; i < numDigesters; i++ {
		g.Go(func() error {
			for path := range paths { // 遍历直到paths chan被关闭
				//...... // 计算path的md5值，放入到c中
				data, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				select {
				case c <- result{path, md5.Sum(data)}:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			return nil
		})
	}
	go func() {
		g.Wait() // 20个goroutine以及遍历文件的goroutine都执行完
		close(c) // 关闭收集结果的chan
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c { // 将md5结果从chan中读取到map中,直到c被关闭才退出
		m[r.path] = r.sum
	}

	// 再次调用Wait，依然可以得到group的error信息
	// Check whether any of the goroutines failed. Since g is accumulating the
	// errors, we don't need to send them (or check for them) in the individual
	// results sent on the channel.
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return m, nil
}

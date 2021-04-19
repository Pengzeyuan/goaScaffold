package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/singleflight"
)

var (
	count = int64(0)
	group = singleflight.Group{}
)

func mainflightDemo() { //flightDemo
	key := "flight"
	for i := 0; i < 5; i++ {
		log.Printf("ID: %d 请求获取缓存", i)
		go func(id int) {
			value, _ := getCache(key, id)
			log.Printf("ID :%d 获取到缓存 , key: %v,value: %v", id, key, value)
		}(i)
	}
	time.Sleep(10 * time.Second)
}

func getCache(key string, id int) (string, error) {
	var ret, _, _ = group.Do(key, func() (ret interface{}, err error) {
		time.Sleep(2 * time.Second) //模拟获取缓存
		log.Printf("ID: %v 执行获取缓存", id)
		return id, nil
	})
	return strconv.Itoa(ret.(int)), nil
}

func main2() { //cyclicBarrierDemo2
	b := cyclicbarrier.NewWithAction(3, func() error {
		fmt.Println("放行")
		return nil
	})

	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("协程 %d 准备好了 \n", id)
			time.Sleep(2 * time.Second)
			// 阻塞等待其它人
			b.Await(context.Background())
			fmt.Printf("协程 %d 通过栅栏 \n", id)
		}(i)
	}

	//b.Await(context.Background())
	//b.Await(context.Background())
	//b.Await(context.Background())
	fmt.Printf("栅栏后方法\n")
	time.Sleep(5 * time.Second)
	log.Printf("完成")
}

func main() { //SingleFight
	g := singleflight.Group{}

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			val, err, shared := g.Do("a", a)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("index: %d, val: %d, shared: %v\n", j, val, shared)
		}(i)
		time.Sleep(time.Millisecond * 100) // 执行时间
	}

	wg.Wait()

}

// 模拟接口方法
func a() (interface{}, error) {
	time.Sleep(time.Duration(rand.Int31n(200)) * time.Millisecond) // 执行时间
	countCur := atomic.AddInt64(&count, 1)
	return countCur, nil
}

//// 部分输出，shared表示是否共享了其他请求的返回结果
//index: 2, val: 1, shared: false
//index: 71, val: 1, shared: true
//index: 69, val: 1, shared: true
//index: 73, val: 1, shared: true
//index: 8, val: 1, shared: true
//index: 24, val: 1, shared: true

//func (g *Group) doCall(c *call, key string, fn func() (interface{}, error)) {
//	c.val, c.err = fn()
//	c.wg.Done()
//
//	g.mu.Lock()
//	if !c.forgotten { // 已调用完，删除这个key
//		delete(g.m, key)
//	}
//	for _, ch := range c.chans {
//		ch <- Result{c.val, c.err, c.dups > 0}
//	}
//	g.mu.Unlock()
//}
//
//// 代表一个正在处理的请求，或者已经处理完的请求
//type call struct {
//	wg sync.WaitGroup
//
//	// 这个字段代表处理完的值，在waitgroup完成之前只会写一次
//	// waitgroup完成之后就读取这个值
//	val interface{}
//	err error
//
//	// 指示当call在处理时是否要忘掉这个key
//	forgotten bool
//	dups      int
//	chans     []chan<- Result
//}
//
//// group代表一个singleflight对象
//type Group struct {
//	mu sync.Mutex       // protects m
//	m  map[string]*call // lazily initialized
//}
//
//func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
//	g.mu.Lock()
//	if g.m == nil {
//		g.m = make(map[string]*call)
//	}
//	if c, ok := g.m[key]; ok { //如果已经存在相同的key
//		c.dups++
//		g.mu.Unlock()
//		c.wg.Wait()               //等待这个key的第一个请求完成
//		return c.val, c.err, true //使用第一个key的请求结果
//	}
//	c := new(call) // 第一个请求，创建一个call
//	c.wg.Add(1)
//	g.m[key] = c //加入到key map中
//	g.mu.Unlock()
//
//	g.doCall(c, key, fn) // 调用方法
//	return c.val, c.err, c.dups > 0
//}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter Counter
	for i := 0; i < 10; i++ { // 10个reader
		go func() {
			for {
				fmt.Println("读操作：", counter.Count()) // 计数器读操作
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for { // 一个writer
		counter.Incr() // 计数器写操作
		time.Sleep(time.Second)
	}
}

// 一个线程安全的计数器
type Counter struct {
	mu    sync.RWMutex
	count uint64
}

// 使用写锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 使用读锁保护
func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

//首先，我们看一下移除了 race 等无关紧要的代码后的 RLock 和 RUnlock 方法：
//func (rw *RWMutex) RLock() {
//	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
//		// rw.readerCount是负值的时候，意味着此时有writer等待请求锁，因为writer优先级高，所以把后来的reader阻塞休眠
//		runtime_SemacquireMutex(&rw.readerSem, false, 0)
//	}
//}
//func (rw *RWMutex) RUnlock() {
//	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
//		rw.rUnlockSlow(r) // 有等待的writer
//	}
//}
//func (rw *RWMutex) rUnlockSlow(r int32) {
//	if atomic.AddInt32(&rw.readerWait, -1) == 0 {
//		// 最后一个reader了，writer终于有机会获得锁了
//		runtime_Semrelease(&rw.writerSem, false, 1)
//	}
//}

//factoria 方法是一个递归计算阶乘的方法，我们用它来模拟 reader。为了更容易地制造出死锁场景，我在这里加上了 sleep 的调用，延缓逻辑的执行。
//这个方法会调用读锁（第 27 行），在第 33 行递归地调用此方法，每次调用都会产生一次读锁的调用，
//所以可以不断地产生读锁的调用，而且必须等到新请求的读锁释放，这个读锁才能释放。

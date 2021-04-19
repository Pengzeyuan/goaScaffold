package mysema

import (
	"container/list"
	"context"
	"sync"
	"unsafe"
)

type Weighted struct {
	size    int64      // 最大资源数
	cur     int64      // 当前已被使用的资源
	mu      sync.Mutex // 互斥锁，对字段的保护
	waiters list.List  // 等待队列
}

// 批量获取资源
func (s *Weighted) Acquire(ctx context.Context, n int64) error {
	s.mu.Lock()
	// fast path, 如果有足够的资源，都不考虑ctx.Done的状态，将cur加上n就返回
	if s.size-s.cur >= n && s.waiters.Len() == 0 {
		s.cur += n
		s.mu.Unlock()
		return nil
	}

	// 如果是不可能完成的任务，请求的资源数大于能提供的最大的资源数
	if n > s.size {
		s.mu.Unlock()
		// 依赖ctx的状态返回，否则一直等待
		<-ctx.Done()
		return ctx.Err()
	}

	// 否则就需要把调用者加入到等待队列中
	// 创建了一个ready chan,以便被通知唤醒
	ready := make(chan struct{})
	w := waiter{n: n, ready: ready} // 等待 一组
	elem := s.waiters.PushBack(w)   // 放到队列尾  s是信号量struct   返回element  e
	s.mu.Unlock()

	// 等待
	select {
	case <-ctx.Done(): //  context的Done被关闭
		err := ctx.Err()
		s.mu.Lock() //  mu是锁
		select {
		case <-ready: // 如果被唤醒了，忽略ctx的状态   // ready slice
			err = nil
		default:
			//通知waiter   返回   第一个 elem  等于插入的  elem
			isFront := s.waiters.Front() == elem // 是list第一个
			s.waiters.Remove(elem)               // 移除 list中的 elem 如果存在
			// 通知其它的waiters,检查是否有足够的资源
			if isFront && s.size > s.cur {
				s.notifyWaiters()
			}
		}
		s.mu.Unlock()
		return err
	case <-ready: // 被唤醒了
		return nil
	}
}

func (s *Weighted) notifyWaiters() {
	for {
		next := s.waiters.Front()
		if next == nil {
			break // No more waiters blocked.
		}

		w := next.Value.(waiter)
		if s.size-s.cur < w.n {
			//避免饥饿，这里还是按照先入先出的方式处理
			break
		}

		s.cur += w.n
		s.waiters.Remove(next)
		close(w.ready)
	}
}

func makechan(t *chantype, size int) *hchan {
	elem := t.elem // chan 类型

	// 略去检查代码
	mem, overflow := math.MulUintptr(elem.size, uintptr(size))

	//
	var c *hchan
	switch {
	case mem == 0:
		// chan的size或者元素的size是0，不必创建buf
		c = (*hchan)(mallocgc(hchanSize, nil, true)) //分配内存
		c.buf = c.raceaddr()                         // c 就是 hchan
	case elem.ptrdata == 0: //
		// 元素不是指针，分配一块连续的内存给hchan数据结构和buf
		c = (*hchan)(mallocgc(hchanSize+mem, nil, true))
		// hchan数据结构后面紧接着就是buf
		c.buf = add(unsafe.Pointer(c), hchanSize) //
	default:
		// 元素包含指针，那么单独分配buf
		c = new(hchan)
		c.buf = mallocgc(mem, elem, true)
	}

	// 元素大小、类型、容量都记录下来
	c.elemsize = uint16(elem.size)
	c.elemtype = elem
	c.dataqsiz = uint(size)
	lockInit(&c.lock, lockRankHchan)

	return c
}

func chansend1(c *hchan, elem unsafe.Pointer) {
	chansend(c, elem, true, getcallerpc())
}
func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
	// 第一部分
	if c == nil {
		if !block {
			return false
		}
		gopark(nil, nil, waitReasonChanSendNilChan, traceEvGoStop, 2)
		throw("unreachable")  11
	}
	......
}
最开始，第一部分是进行判断：如果 chan 是 nil 的话，就把调用者 goroutine park（阻塞休眠）， 调用者就永远被阻塞住了，所以，第 11 行是不可能执行到的代码。
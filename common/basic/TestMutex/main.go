package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false) // 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	fmt.Println("id:", id)
	return id
}

func test9() {

}

var num = 3

//  并发 Cond
type FIFO struct {
	lock  sync.Mutex
	cond  *sync.Cond
	queue []int
}

type Queue interface {
	Pop() int
	Offer(num int) error
}

func (f *FIFO) Offer(num int) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.queue = append(f.queue, num)
	f.cond.Broadcast()
	return nil
}
func (f *FIFO) Pop() int {
	f.lock.Lock()
	defer f.lock.Unlock()
	for {
		for len(f.queue) == 0 {
			f.cond.Wait()
		}
		item := f.queue[0]
		f.queue = f.queue[1:]
		return item
	}
}

func main() {
	GoID()
	//testMainMutexCond()
}

func testMainMutexCond() {
	l := sync.Mutex{}
	fifo := &FIFO{
		lock:  l,
		cond:  sync.NewCond(&l),
		queue: []int{},
	}
	go func() {
		for {
			fifo.Offer(rand.Int())
		}
	}()
	time.Sleep(time.Second)
	go func() {
		for {
			fmt.Println(fmt.Sprintf("goroutine1 pop-->%d", fifo.Pop()))
		}
	}()
	go func() {
		for {
			fmt.Println(fmt.Sprintf("goroutine2 pop-->%d", fifo.Pop()))
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

//  优雅的关闭  goroutine
type Service struct {
	// Other things

	ch        chan bool
	waitGroup *sync.WaitGroup
}

func NewService() *Service {
	s := &Service{
		// Init Other things
		ch:        make(chan bool),
		waitGroup: &sync.WaitGroup{},
	}

	return s
}

func (s *Service) Stop() {
	close(s.ch)
	s.waitGroup.Wait()
}

func (s *Service) Serve() {
	s.waitGroup.Add(1)
	defer s.waitGroup.Done()

	for {
		select {
		case <-s.ch:
			fmt.Println("stopping...")
			return
		default:
			fmt.Println("doing something")
		}
		s.waitGroup.Add(1)
		go s.anotherServer()
	}
}
func (s *Service) anotherServer() {
	defer s.waitGroup.Done()
	for {
		select {
		case <-s.ch:
			fmt.Println("stopping...")
			return
		default:
			fmt.Println("doing otherThing")
		}

		// Do something
	}
}

//func main() {
//
//	service := NewService()
//	go service.Serve()
//
//	// Handle SIGINT and SIGTERM.
//	ch := make(chan os.Signal)
//	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
//	fmt.Println(<-ch)
//
//	// Stop the service gracefully.
//	service.Stop()
//}

//func main() {
//	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
//	for i := 1; i <= int(math.Floor(float64(len(ints)/num)))+1; i++ {
//		low := num * (i - 1)
//		if low > len(ints) {
//			return
//		}
//		high := num * i
//		if high > len(ints) {
//			high = len(ints)
//		}
//		glog.Info(ints[low:high])
//	}
//}

package main

import (
	"fmt"

	"time"
)

func Test2IntFloat() {
	var duration_Milliseconds time.Duration = 500 * time.Millisecond

	var castToInt64 int64 = duration_Milliseconds.Nanoseconds() / 1e6
	var castToFloat64 float64 = duration_Milliseconds.Seconds() * 1e3
	fmt.Printf("Duration [%v]\ncastToInt64 [%d]\ncastToFloat64 [%.0f]\n", duration_Milliseconds, castToInt64, castToFloat64)
}
func TestdurationDseconds() {
	var duration_Seconds time.Duration = (1250 * 10) * time.Millisecond
	var duration_Minute time.Duration = 2 * time.Minute

	var float64_Seconds float64 = duration_Seconds.Seconds()
	var float64_Minutes float64 = duration_Minute.Minutes()

	fmt.Printf("Seconds [%.3f]\nMinutes [%.2f]\n", float64_Seconds, float64_Minutes)
}

func TestTime() {
	TestdurationDseconds()
	testDUration()

	var waitFiveHundredMillisections time.Duration = 500 * time.Millisecond

	startingTime := time.Now().UTC()
	time.Sleep(600 * time.Millisecond)
	endingTime := time.Now().UTC()
	//  就是真正持续的时间
	var duration time.Duration = endingTime.Sub(startingTime)

	if duration >= waitFiveHundredMillisections {
		fmt.Printf("Wait %v\nNative [%v]\nMilliseconds [%d]\nSeconds [%.3f]\n", waitFiveHundredMillisections, duration, duration.Nanoseconds()/1e6, duration.Seconds())
	}

}

func testDUration() {
	var duration_Milliseconds time.Duration = 500 * time.Millisecond
	var duration_Seconds time.Duration = (1250 * 10) * time.Millisecond
	var duration_Minute time.Duration = 2 * time.Minute

	fmt.Printf("Milli [%v]\nSeconds [%v]\nMinute [%v]\n", duration_Milliseconds, duration_Seconds, duration_Minute)
}

func main() {
	TestTime()
	//TimeAfter()
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

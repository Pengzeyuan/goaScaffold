package main

import (
	"github.com/golang/glog"
	"math"
)

var num = 3

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	for i := 1; i <= int(math.Floor(float64(len(ints)/num)))+1; i++ {
		low := num * (i - 1)
		if low > len(ints) {
			return
		}
		high := num * i
		if high > len(ints) {
			high = len(ints)
		}
		glog.Info(ints[low:high])
	}
}

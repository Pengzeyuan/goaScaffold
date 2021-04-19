package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Log("hello world")
}

func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

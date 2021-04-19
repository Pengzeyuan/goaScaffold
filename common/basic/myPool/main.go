package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"sync"
)

var buffers = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}

func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	buffers.Put(buf)
}

type channelPool struct { // 存储连接池的
	mu    sync.RWMutex
	conns chan net.Conn // net.Conn 的产生器
	//factory Factory

}

var pool *sync.Pool

type Person struct {
	Name string
}

func init() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("creating a new person")
			return new(Person)
		},
	}
}

func fnPool() {

	person := pool.Get().(*Person)
	fmt.Println("Get Pool Object：", person)

	person.Name = "first"
	pool.Put(person)

	fmt.Println("Get Pool Object：", pool.Get().(*Person))
	fmt.Println("Get Pool Object：", pool.Get().(*Person))

}

func main() {
	background := context.Background()
	fnPool()

	//byteContent := []byte("abcd")

	//getedbuffer := buffers.Get()

	//fmt.Printf("%t,", getedbuffer)
	//getedbuffer.Write(byteContent)
	buffers = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	b := bytes.NewBufferString("abcde")
	PutBuffer(b)
	getBuffer := GetBuffer()
	fmt.Println("jieguo:", getBuffer)
	fmt.Println("pool:", buffers)

}

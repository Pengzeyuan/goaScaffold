package websocket

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Connection struct {
	wsConnect *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte

	mutex    sync.Mutex // 对closeChan关闭上锁
	isClosed bool       // 防止closeChan被关闭多次
}

func NewConnection(wsConn *websocket.Conn) (conn *Connection) {
	conn = &Connection{
		wsConnect: wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}
	return
}

// 内部实现
func (conn *Connection) ReadLoop() {
	var (
		data []byte
		err  error
	)
	for {
		// 底层的read message
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			goto ERR
		}
		if data != nil {
			// 写 message  把pong 写给 outChan
			_ = conn.WriteMessage([]byte("pong"))
		}
		//阻塞在这里，等待inChan有空闲位置
		select {
		//case conn.inChan <- data: // 因为不需要消费，防止堵塞，所以取消
		case <-conn.closeChan: // closeChan 感知 conn断开
			goto ERR
		case <-time.After(time.Millisecond * 500): // 若没有能执行的，执行超时并让出cpu执行权
		}
	}

ERR:
	// 关闭 socket
	conn.Close()
}

//  取出  outChan  的data
func (conn *Connection) WriteLoop() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		//  底层的写入  把data
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}
func (conn *Connection) WriteMessage(data []byte) (err error) {
	//  写入  outChan
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closeed")
	}
	return
}

func (conn *Connection) Close() {
	// 线程安全，可多次调用  底层
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.isClosed {
		// 关闭 chan
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

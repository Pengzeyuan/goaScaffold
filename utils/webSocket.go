package utils

import (
	"boot/config"
	"boot/model"

	"encoding/json"

	"fmt"

	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"

	"go.uber.org/zap"

	"net/http"

	"reflect"

	"strings"

	"time"
)

var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// websocket 列表
	connList []*Connection
	//  nat 消息体  传输列表
	natsDataChan = make(chan *model.NatsData)
)

func SetOff(logger *zap.Logger) { // 检测离开连接
	var (
		err error
	)
	for {
		for k := 0; k < len(connList); k++ {
			thisConn := connList[k]
			if err = thisConn.WriteMessage([]byte("Ping")); err != nil {
				thisConn.Close()
				connList = append(connList[:k], connList[k+1:]...)
				logger.Info(fmt.Sprintf("有 1 个连接离开了,当前连接人数: %d", len(connList)))
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func (conn *Connection) Close() {
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// websocket
func ServerWs(logger *zap.Logger, w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *Connection
		data   []byte
	)

	//subscribeUserID := r.FormValue("userId") // 可从前端获取用户id
	//subscribeUserID := "userId"
	// 完成ws协议的握手操作
	// Upgrade:websocket
	logger.Info("websocket开始建立连接")

	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil { // 升级为websocket
		logger.Error("连接 websocket 失败", zap.Error(err))
		return
	}

	//  把 websocket 在封装成 connection
	conn = NewConnection(wsConn) // 初始化conn
	connList = append(connList, conn)
	logger.Info(fmt.Sprintf("新来了 1 个连接,当前连接人数: %d", len(connList)))

	// 这里是  收  数据库的某行变化得到结构体
	_, _ = NatsCli.QueueSubscribe(config.C.Nats.CanalSubject, "queue", func(m *nats.Msg) {
		logger.Info(fmt.Sprintf("nat 接收到的消息data: %d", m.Data))
		logger.Info(fmt.Sprintf("map的值 %s %s \n", m.Data, reflect.TypeOf(m.Data)))

		//转换字符串 为json string
		sprintf := string(m.Data)

		str := strings.Replace(sprintf, "\n", "", -1)
		//str = "\"" + str + "\""
		logger.Info(fmt.Sprintln("jsonString:", str))
	})
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()

	for {
		//  nats 订阅
		_, err = NatsCli.QueueSubscribe(config.C.Nats.CallNoSubject, "queue", func(m *nats.Msg) {
			natsMessage := &model.PushInfoEvent{}
			if err = json.Unmarshal(m.Data, &natsMessage); err != nil {
				logger.Error("反序列化失败出错", zap.Error(err))
				return
			}
			//logger.Info("订阅 nats 消息成功")
			natsData := model.NatsData{}
			natsData.CallNoID = natsMessage.CallNoID
			natsData.DataBytes = m.Data
			// 放入传输单例
			natsDataChan <- &natsData
		})
		if err != nil {
			logger.Error("订阅 websocket 通知失败", zap.Error(err))
			continue
		}
		_ = NatsCli.Flush()
		//  取出一条 nat消息
		natsData := <-natsDataChan
		//  natsMessage.byte
		data = natsData.DataBytes
		logger.Info(fmt.Sprintln("获得的data消息", string(data)))
		for i := 0; i < len(connList); i++ {
			sendConn := connList[i]
			// 其中一个  websocket   写入数据到outChan
			if err = sendConn.WriteMessage(data); err != nil {
				logger.Error("发送 websocket 消息失败", zap.Error(err))
				// websocket 关闭
				sendConn.Close()
			} else {
				logger.Info("发送 websocket 消息成功")
			}
		}
	}
}

// 内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		//  websocket 读出数据  【】byte conn连接中的读者器中的全部信息
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			goto ERR
		}
		//阻塞在这里，等待inChan有空闲位置
		//data 存入 inChan
		select {
		case conn.inChan <- data:
		case <-conn.closeChan: // closeChan 感知 conn断开
			goto ERR
		}

	}

ERR:
	conn.Close()
}

// 写循环  先取 outCHan的内容
func (conn *Connection) writeLoop() {
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
		// 写入数据 到  conn 中
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}

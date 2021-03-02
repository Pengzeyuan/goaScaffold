package websocket

import (
	"go.uber.org/zap"
	"sync"
)

type Manager struct {
	lock     sync.Mutex
	handlers map[string]*Connection
}

func NewManager() *Manager {
	return &Manager{
		handlers: make(map[string]*Connection),
		lock:     sync.Mutex{},
	}
}

// 订阅
func (m *Manager) Subscribe(userID string, conn *Connection) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	_, ok := m.handlers[userID]
	if !ok {
		m.handlers[userID] = conn
	}
	return nil

}

// 发布
func (m *Manager) Publish(data []byte, logger *zap.Logger) {
	for _, v := range m.handlers {
		//  子写的  写入数据到  outchan
		if err := v.WriteMessage(data); err != nil {
			logger.Error("发送 websocket 消息失败", zap.Error(err))
			v.Close()
		}
	}
}

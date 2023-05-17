package ws

import (
	"sync"
)

// Manager 连接管理模块
type Manager struct {
	mu sync.RWMutex

	connections map[uint64]Connection
}

// NewManager 创建一个链接管理器
func NewManager() *Manager {
	return &Manager{
		connections: make(map[uint64]Connection),
	}
}

// Add 添加连接
func (c *Manager) Add(conn Connection) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connections[conn.GetID()] = conn
}

// Remove 删除连接
func (c *Manager) Remove(conn Connection) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.connections, conn.GetID())
}

// Get 获取连接
func (c *Manager) Get(cid uint64) (Connection, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if conn, ok := c.connections[cid]; ok {
		return conn, nil
	}
	return nil, ErrConnNotFound
}

// Len 获取连接数
func (c *Manager) Len() int {
	return len(c.connections)
}

// Clear 清除并停止所有连接
func (c *Manager) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for cid, conn := range c.connections {
		go conn.Stop()
		delete(c.connections, cid)
	}
}

// Range 遍历所有连接
func (c *Manager) Range(f ConnHandlerFunc) (err error) {
	for cid, conn := range c.connections {
		err = f(cid, conn)
	}

	return err
}

package ws

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/binbinly/pkg/logger"
	"github.com/gorilla/websocket"
)

// Connection 定义连接接口
type Connection interface {
	Start()                   //启动连接，让当前连接开始工作
	Stop()                    //停止连接，结束当前连接状态
	Context() context.Context //返回ctx，用于用户自定义的go程获取连接退出状态

	GetID() uint64        //获取当前连接ID
	GetUID() int          //获取当前连接鉴权ID
	RemoteAddr() net.Addr //获取远程客户端地址信息

	Send(ctx context.Context, mid int, data []byte) error      //发送消息
	AsyncSend(ctx context.Context, mid int, data []byte) error //异步发送消息
}

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

// wsConnection 连接
type wsConnection struct {
	mu sync.RWMutex
	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	id  uint64
	uid int
	//告知该链接已经退出/停止的channel
	ctx    context.Context
	cancel context.CancelFunc
	//消息管道，用于读、写两个goroutine之间的消息通信
	msgChan chan []byte
	//当前Conn属于哪个Server
	server *wsServer
	//websocket连接对象
	conn *websocket.Conn
	//websocket响应对象
	writer http.ResponseWriter
	//websocket请求对象
	request *http.Request
	//当前连接的关闭状态
	isClosed bool
}

// NewConnect 创建连接的方法
func NewConnect(s *wsServer, conn *websocket.Conn, id uint64, uid int) Connection {
	return &wsConnection{
		server:  s,
		id:      id,
		uid:     uid,
		conn:    conn,
		msgChan: make(chan []byte, s.Options().MaxMsgChanLen),
	}
}

// startWriter 写消息Goroutine， 用户将数据发送给客户端
func (c *wsConnection) startWriter() {
	logger.Debug("[ws.write] Writer Goroutine is running")
	// 心跳由客户端发送 ping 回复 pong
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		defer logger.Debugf("[ws.write] %v conn Writer exit!", c.RemoteAddr().String())
		ticker.Stop()
	}()

	for {
		select {
		case data, ok := <-c.msgChan:
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte("close"))
				return
			}
			_ = c.conn.SetWriteDeadline(time.Now().Add(c.server.Options().WriteWait))
			logger.Debugf("[ws.write] write msg:%v", string(data))
			if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				logger.Warnf("[ws.write] conn.write err :%v", err)
				break
			}
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(c.server.Options().WriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Warnf("[ws.write] ticker ping wrr:%v", err)
				return
			}
		}
	}
}

// startReader 读消息Goroutine，用于从客户端中读取数据
func (c *wsConnection) startReader() {
	logger.Debug("[ws.read] Reader Goroutine is running")
	defer c.Stop()

	c.conn.SetReadLimit(int64(c.server.Options().MaxPacketSize))
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			mType, buff, err := c.conn.ReadMessage()
			if err != nil {
				logger.Infof("[ws.read] readPump ReadMessage err:%V", err)
				c.cancel()
				return
			}
			if mType == websocket.PingMessage {
				continue
			}
			if len(buff) == 0 {
				continue
			}
			logger.Debugf("[ws.read] reader message:%v", string(buff))
			if string(buff) == "ping" {
				_ = c.AsyncSend(context.Background(), 0, []byte("pong"))
			} else {
				// 构建当前客户端请求的request数据
				req, err := NewRequest(c, buff)
				if err != nil {
					logger.Warnf("[ws.read] data format err:%v", err)
					return
				}
				c.server.handler.AsyncExecute(req)
			}
		}
	}
}

// Start 启动连接，让当前连接开始工作
func (c *wsConnection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())

	if c.server.Options().OnConnStart != nil {
		c.server.Options().OnConnStart(c)
	}

	// 开启用户从客户端读取数据的Goroutine
	go c.startReader()
	// 开启用于写回客户端数据的Goroutine
	go c.startWriter()
}

// Stop 关闭连接
func (c *wsConnection) Stop() {
	//如果用户注册了该连接的关闭回调业务，那么在此调用
	if c.server.Options().OnConnStop != nil {
		c.server.Options().OnConnStop(c)
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}

	//关闭socket连接
	if err := c.conn.Close(); err != nil {
		logger.Warnf("[ws.stop] connection closed err:%v", err)
	}
	//关闭Writer
	c.cancel()

	// 将连接从连接管理器中删除
	c.server.GetManager(c.id).Remove(c)
	// 关闭该连接全部管道
	close(c.msgChan)
	//设置标志位
	c.isClosed = true
}

func (c *wsConnection) Context() context.Context {
	return c.ctx
}

// GetID 获取连接id
func (c *wsConnection) GetID() uint64 {
	return c.id
}

// GetUID 获取连接id
func (c *wsConnection) GetUID() int {
	return c.uid
}

// RemoteAddr 获取远程客户端地址信息
func (c *wsConnection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// Send 发送数据给远程的WS客户端
func (c *wsConnection) Send(ctx context.Context, mid int, msg []byte) error {
	if c.isClosed == true {
		return ErrConnNotFinish
	}

	//写回客户端
	if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return err
	}
	return nil
}

// AsyncSend 异步发送数据给远程的WS客户端
func (c *wsConnection) AsyncSend(ctx context.Context, mid int, msg []byte) error {
	if c.isClosed == true {
		return ErrConnNotFinish
	}

	//写回客户端
	c.msgChan <- msg
	return nil
}

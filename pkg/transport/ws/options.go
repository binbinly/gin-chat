package ws

import (
	"net/http"
	"time"
)

var defOptions = &Options{
	Name:             "WebSocket",
	Addr:             ":9060",
	MaxPacketSize:    4096,
	MaxConn:          10000,
	WorkerPoolSize:   1,
	MaxWorkerTaskLen: 128,
	MaxMsgChanLen:    128,
	ManagerSize:      1,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	WriteWait:        10 * time.Second,
}

type AuthHandler = func(r *http.Request, sid string, cid uint64) (uid int, ok bool)

type Option func(*Options)

type Options struct {
	ID               string        //服务器ID
	Name             string        //服务器的名称
	Addr             string        //服务绑定的地址
	MaxPacketSize    int           //都需数据包的最大值
	MaxConn          int           //当前服务器主机允许的最大链接个数
	WorkerPoolSize   int           //业务工作Worker池的数量
	MaxWorkerTaskLen int           //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    int           //SendBuffMsg发送消息的缓冲最大长度
	ManagerSize      int           //连接管理器个数
	ReadBufferSize   int           //接收缓冲区
	WriteBufferSize  int           //发送缓冲区
	WriteWait        time.Duration //写入客户端超时

	Router      *Engine               //请求路由
	OnConnStart func(conn Connection) //该Server的连接创建开始时Hook函数
	OnConnStop  func(conn Connection) //该Server的连接断开时的Hook函数
	OnConnAuth  AuthHandler           //该Server的连接鉴权完成的Hook函数
}

func WithID(id string) Option {
	return func(o *Options) {
		o.ID = id
	}
}

func WithName(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

func WithAddr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}

func WithRouter(r *Engine) Option {
	return func(o *Options) {
		o.Router = r
	}
}

func WithMaxPacketSize(size int) Option {
	return func(o *Options) {
		o.MaxPacketSize = size
	}
}

func WithMaxConn(size int) Option {
	return func(o *Options) {
		o.MaxConn = size
	}
}

func WithWorkerPoolSize(size int) Option {
	return func(o *Options) {
		o.WorkerPoolSize = size
	}
}

func WithMaxWorkerTaskLen(size int) Option {
	return func(o *Options) {
		o.MaxWorkerTaskLen = size
	}
}

func WithMaxMsgChanLen(size int) Option {
	return func(o *Options) {
		o.MaxMsgChanLen = size
	}
}

func WithManagerSize(size int) Option {
	return func(o *Options) {
		o.ManagerSize = size
	}
}

func WithReadBufferSize(size int) Option {
	return func(o *Options) {
		o.ReadBufferSize = size
	}
}

func WithWriteBufferSize(size int) Option {
	return func(o *Options) {
		o.WriteBufferSize = size
	}
}

func WithWriteWait(d time.Duration) Option {
	return func(o *Options) {
		o.WriteWait = d
	}
}

func WithOnConnStart(f func(conn Connection)) Option {
	return func(o *Options) {
		o.OnConnStart = f
	}
}

func WithOnConnStop(f func(conn Connection)) Option {
	return func(o *Options) {
		o.OnConnStop = f
	}
}

func WithOnConnAuth(f AuthHandler) Option {
	return func(o *Options) {
		o.OnConnAuth = f
	}
}

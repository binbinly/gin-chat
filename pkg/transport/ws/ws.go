package ws

import (
	"context"
	"errors"
	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/util"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"github.com/zhenjl/cityhash"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

var (
	// ErrConnNotFound 连接未找到
	ErrConnNotFound = errors.New("connection not found")
	// ErrConnNotFinish 连接未完成，不可以发送消息
	ErrConnNotFinish = errors.New("connection not finish when send msg")
)

type ConnHandlerFunc func(cid uint64, conn Connection) error

// Server is a simple micro server abstraction
type Server interface {
	// Init Initialise options
	Init(...Option)
	// Options Retrieve the options
	Options() *Options
	// Start the server
	Start(ctx context.Context) error
	// Stop the server
	Stop(ctx context.Context) error
	// Endpoint return a real address to registry endpoint.
	Endpoint() (*url.URL, error)
	// GetManager 所有连接管理
	GetManager(cid uint64) *Manager
	// Range 遍历所有连接
	Range(f ConnHandlerFunc)
	// Total 服务器连接总数
	Total() int
}

// wsServer 基础服务
type wsServer struct {
	managers []*Manager
	handler  *Handler
	opts     *Options
	lis      net.Listener
	endpoint *url.URL
	upgrader *websocket.Upgrader
}

// NewServer 实例化websocket服务器
func NewServer() Server {
	return &wsServer{
		opts: defOptions,
	}
}

// Options 服务选项
func (s *wsServer) Options() *Options {
	return s.opts
}

// Init 初始化
func (s *wsServer) Init(opts ...Option) {
	for _, o := range opts {
		o(s.opts)
	}
	if s.opts.ID == "" {
		s.opts.ID = xid.New().String()
	}
	//初始化连接管理器
	s.managers = make([]*Manager, s.opts.ManagerSize)
	for i := 0; i < s.opts.ManagerSize; i++ {
		s.managers[i] = NewManager()
	}
	//初始化消息处理器
	s.handler = NewHandler(s.opts.WorkerPoolSize, s.opts.Router)
	s.upgrader = &websocket.Upgrader{
		ReadBufferSize:  s.opts.ReadBufferSize,
		WriteBufferSize: s.opts.WriteBufferSize,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
}

// Start 启动服务器
func (s *wsServer) Start(ctx context.Context) error {
	// 启动worker工作池机制
	s.handler.Init(s.opts.MaxWorkerTaskLen)
	return s.Listen()
}

// Stop 关闭服务器
func (s *wsServer) Stop(ctx context.Context) error {
	log.Print("[Websocket] server is stopping")
	for _, manager := range s.managers {
		manager.Clear()
	}
	return s.lis.Close()
}

// Listen websocket连接监听
func (s *wsServer) Listen() error {
	var cid uint64
	lis, err := net.Listen("tcp", s.opts.Addr)
	if err != nil {
		return err
	}
	s.lis = lis

	if _, err = s.Endpoint(); err != nil {
		return err
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//设置服务器最大连接控制,如果超过最大连接，则拒绝
		if s.Total() >= s.Options().MaxConn {
			logger.Warn("[ws.start] connection size limit")
			return
		}
		// 如果需要 websocket 认证请设置认证信息
		uid := 0
		if s.Options().OnConnAuth != nil {
			var ok bool
			if uid, ok = s.Options().OnConnAuth(r, s.opts.ID, cid); !ok {
				w.WriteHeader(401)
				return
			}
		}
		// 判断 header 里面是有子协议
		if len(r.Header.Get("Sec-Websocket-Protocol")) > 0 {
			s.upgrader.Subprotocols = websocket.Subprotocols(r)
		}
		// 升级成 websocket 连接
		c, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		conn := NewConnect(s, c, cid, uid)
		// 添加连接至管理器
		s.GetManager(cid).Add(conn)
		conn.Start()
		cid++
	})
	log.Printf("[Websocket] server is listening on: %s", lis.Addr().String())
	if err = http.Serve(lis, nil); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

// Endpoint return a real address to registry endpoint.
// examples: http://127.0.0.1:8080
func (s *wsServer) Endpoint() (*url.URL, error) {
	addr, err := util.Extract(s.opts.Addr, s.lis)
	if err != nil {
		return nil, err
	}
	s.endpoint = &url.URL{Scheme: "http", Host: addr}
	return s.endpoint, nil
}

// GetManager 获取当前连接的管理器
func (s *wsServer) GetManager(cid uint64) *Manager {
	str := strconv.FormatUint(cid, 10)
	idx := cityhash.CityHash32([]byte(str), uint32(len(str))) % uint32(s.opts.ManagerSize)
	return s.managers[idx]
}

// Range 遍历所有连接
func (s *wsServer) Range(f ConnHandlerFunc) {
	for _, manager := range s.managers {
		_ = manager.Range(f)
	}
}

// Total 当前服务器的总连接数
func (s *wsServer) Total() int {
	var c int
	for _, manager := range s.managers {
		c += manager.Len()
	}
	return c
}

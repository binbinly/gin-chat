package ws

import (
	"math"
	"sync"
	"time"
)

const abortIndex int8 = math.MaxInt8 / 2 // 路由中间件嵌套最大值

// HandlerFunc 路由处理方法
type HandlerFunc func(c *Context)

// HandlerChain 执行链
type HandlerChain []HandlerFunc

// RouterGroup 路由组
type RouterGroup struct {
	Handlers HandlerChain
	engine   *Engine
}

// Engine 路由引擎
type Engine struct {
	RouterGroup

	tree map[string]HandlerChain
	pool sync.Pool
}

// Context 上下文对象
type Context struct {
	Req      *Request
	handlers HandlerChain
	index    int8
}

// NewEngine get a new engine(tcp.Handler)
func NewEngine() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
		},
		tree: make(map[string]HandlerChain),
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

// Next 执行下一个
func (c *Context) Next() {
	c.index++
	// 调用 c.Abort() 的时候不会往后执行
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// Abort 退出路由
func (c *Context) Abort() {
	c.index = abortIndex
}

// Reset 重置路由
func (c *Context) Reset() {
	c.handlers = nil
	c.index = -1
}

// Done always returns nil (chan which will wait forever),
// if you want to abort your work when the connection was closed
// you should use Request.Context().Done() instead.
func (c *Context) Done() <-chan struct{} {
	return nil
}

// Err always returns nil, maybe you want to use Request.Context().Err() instead.
func (c *Context) Err() error {
	return nil
}

// Deadline always returns that there is no deadline (ok==false),
// maybe you want to use Request.Context().Deadline() instead.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

// Value returns the value associated with this context for key, or nil
// if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
func (c *Context) Value(interface{}) interface{} {
	return c.Req
}

// allocateContext
func (e *Engine) allocateContext() *Context {
	return &Context{}
}

// Start 执行入口
func (e *Engine) Start(req *Request) {
	c := e.pool.Get().(*Context)
	c.Req = req
	c.Reset()
	handlers := e.getHandlers(c.Req)
	if handlers != nil {
		c.handlers = handlers
		c.Next()
	}
}

func (e *Engine) addRoute(event string, handlers HandlerChain) {
	e.tree[event] = handlers
}

func (e *Engine) getHandlers(req *Request) HandlerChain {
	handlers, ok := e.tree[req.Event()]
	if !ok {
		return nil
	}
	return handlers
}

// Use set common middleware
func (e *Engine) Use(middleware ...HandlerFunc) {
	e.RouterGroup.Use(middleware...)
}

// Use 加载中间件
func (g *RouterGroup) Use(middleware ...HandlerFunc) {
	g.Handlers = append(g.Handlers, middleware...)
}

// AddRoute specific middleware
func (g *RouterGroup) AddRoute(event string, handlers ...HandlerFunc) {
	handlers = g.mergeHandlers(handlers)
	g.engine.addRoute(event, handlers)
}

// merge specific and common middleware
func (g *RouterGroup) mergeHandlers(handlers HandlerChain) HandlerChain {
	finalSize := len(g.Handlers) + len(handlers)
	mergedHandlers := make(HandlerChain, finalSize)
	copy(mergedHandlers, g.Handlers)
	copy(mergedHandlers[len(g.Handlers):], handlers)
	return mergedHandlers
}

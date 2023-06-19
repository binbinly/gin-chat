package router

import (
	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/transport/ws"

	"gin-chat/internal/websocket"
)

// NewWsRouter 实例化websocket路由
func NewWsRouter() *ws.Engine {
	r := ws.NewEngine()
	r.Use(func(c *ws.Context) {
		logger.Infof("[ws] event: %v", c.Req.Event())
		c.Next()
	})
	r.AddRoute("ping", Ping)
	return r
}

// Ping 心跳
func Ping(c *ws.Context) {
	if err := c.Req.Conn().Send(c, 0, websocket.Pack("pong", "")); err != nil {
		logger.Info("[ws.ping] err: %v", err)
	}
}

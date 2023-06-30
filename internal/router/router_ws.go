package router

import (
	"gin-chat/internal/service"
	"gin-chat/internal/websocket"

	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/transport/ws"
)

// NewWsRouter 实例化websocket路由
func NewWsRouter() *ws.Engine {
	r := ws.NewEngine()
	r.Use(func(c *ws.Context) {
		logger.Debugf("[ws] event: %v", c.Req.Event())
		c.Next()
	})
	r.AddRoute("ping", Ping)
	r.AddRoute("history", History)

	return r
}

// Ping 心跳
func Ping(c *ws.Context) {
	if err := c.Req.Conn().Send(c, 0, websocket.Pack("pong", "")); err != nil {
		logger.Warnf("[ws.ping] err: %v", err)
	}
}

// History 用户历史
func History(c *ws.Context) {
	if err := service.Svc.PushHistory(c, c.Req.Conn().GetUID()); err != nil {
		logger.Warnf("[ws.history] err: %v", err)
	}
}

package server

import (
	"context"
	"net/http"

	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/transport/ws"
	"github.com/rs/xid"

	"gin-chat/internal/router"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// online test url: http://www.websocket-test.com

// NewWsServer websocket server
func NewWsServer(conf *app.ServerConfig) ws.Server {
	s := ws.NewServer()
	s.Init(ws.WithID(xid.New().String()),
		ws.WithAddr(conf.Addr),
		ws.WithWriteWait(conf.WriteTimeout),
		ws.WithRouter(router.NewWsRouter()),
		ws.WithOnConnAuth(onConnectionAuth()),
		ws.WithOnConnStop(onConnectionLost))
	return s
}

// onConnectionAuth 与客户端建立连接后鉴权
func onConnectionAuth() ws.AuthHandler {
	return func(r *http.Request, sid string, cid uint64) (int, bool) {
		token := r.URL.Query().Get("token")
		if token == "" {
			return 0, false
		}

		uid, err := service.Svc.UserOnline(r.Context(), token, sid, cid)
		if err != nil {
			logger.Debugf("[ws.conn] user online err: %v token: %v", err, token)
			return 0, false
		}
		logger.Debugf("[ws.conn] user online success to %v", uid)
		return uid, true
	}
}

// onConnectionLost 与客户端断开连接时执行
func onConnectionLost(conn ws.Connection) {
	logger.Debug("Do Connection lost is Called ...")
	// 不可以用 conn.Context() 连接可能已经取消 会报：context canceled
	if err := service.Svc.UserOffline(context.Background(), conn.GetUID()); err != nil {
		logger.Warnf("[ws.conn] lost offline err:%v", err)
	}
}

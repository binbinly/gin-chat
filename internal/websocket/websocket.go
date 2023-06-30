package websocket

import (
	"context"
	"encoding/json"

	"gin-chat/pkg/app"

	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/transport/ws"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// UserConnInfo 用户连接信息
type UserConnInfo struct {
	UserID   int    `json:"user_id"`
	ConnID   uint64 `json:"conn_id"`
	ServerID string `json:"server_id"`
}

// Server websocket server
type Server struct {
	ws  ws.Server
	rdb *redis.Client
}

// New websocket server
func New(ws ws.Server, rdb *redis.Client) *Server {
	return &Server{
		ws:  ws,
		rdb: rdb,
	}
}

// Send 发送消息
func (w *Server) Send(ctx context.Context, c *UserConnInfo, event string, data any) (err error) {
	return w.send(ctx, c, Pack(event, data))
}

// BatchSendConn 批量发送多个连接
func (w *Server) BatchSendConn(ctx context.Context, cs []*UserConnInfo, event string, data any) (err error) {
	msg := Pack(event, data)
	for _, c := range cs {
		if err = w.send(ctx, c, msg); err != nil {
			logger.Warnf("[ws.batchSend] err: %v", c.UserID, err)
		}
	}

	return err
}

// BatchSendMessage 批量发送多条消息
func (w *Server) BatchSendMessage(ctx context.Context, c *UserConnInfo, list []string) error {
	if len(list) == 0 {
		return nil
	}
	conn, err := w.ws.GetManager(c.ConnID).Get(c.ConnID)
	if errors.Is(err, ws.ErrConnNotFound) {
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "[ws.Close] get conn by uid: %v", c.UserID)
	}
	for _, msg := range list {
		if err = conn.AsyncSend(ctx, 0, []byte(msg)); err != nil {
			logger.Warnf("[ws.BatchSendMessage] AsyncSend uid:%d err: %v", c.UserID, err)
		}
	}

	return err
}

// Close 发送客户端连接
func (w *Server) Close(ctx context.Context, c *UserConnInfo, data any) error {
	conn, err := w.ws.GetManager(c.ConnID).Get(c.ConnID)
	if errors.Is(err, ws.ErrConnNotFound) {
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "[ws.Close] get conn by uid: %v", c.UserID)
	}

	if data != nil {
		_ = conn.Send(ctx, 0, Pack(EventClose, data))
	}
	conn.Stop()

	return nil
}

// Broadcast 广播消息
func (w *Server) Broadcast(ctx context.Context, msg []byte) (err error) {
	w.ws.Range(func(cid uint64, conn ws.Connection) error {
		return w.send(ctx, &UserConnInfo{ConnID: cid}, msg)
	})

	return
}

// send 发送消息至客户端
func (w *Server) send(ctx context.Context, c *UserConnInfo, msg []byte) error {
	if c.ConnID == 0 {
		w.saveHistory(ctx, c, msg)
		return nil
	}
	conn, err := w.ws.GetManager(c.ConnID).Get(c.ConnID)
	if errors.Is(err, ws.ErrConnNotFound) {
		w.saveHistory(ctx, c, msg)
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "[ws.Close] get conn by uid: %v", c.ConnID)
	}

	if err = conn.AsyncSend(ctx, 0, msg); err != nil {
		return errors.Wrapf(err, "[ws.send] AsyncSend msg:%v", msg)
	}

	return nil
}

func (w *Server) saveHistory(ctx context.Context, c *UserConnInfo, msg []byte) {
	w.rdb.RPush(ctx, app.BuildHistoryKey(c.UserID), msg)
}

// Pack 数据打包 json 传输
func Pack(event string, data any) []byte {
	msg, _ := json.Marshal(struct {
		Event string `json:"event"`
		Data  any    `json:"data"`
	}{
		Event: event,
		Data:  data,
	})
	return msg
}

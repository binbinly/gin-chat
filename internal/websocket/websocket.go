package websocket

import (
	"context"
	"encoding/json"

	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/transport/ws"
	"github.com/pkg/errors"
)

// Server websocket server
type Server struct {
	ws ws.Server
}

// New websocket server
func New(ws ws.Server) *Server {
	return &Server{
		ws: ws,
	}
}

// Send 发送消息
func (w *Server) Send(ctx context.Context, id uint64, event string, data any) (err error) {
	return w.send(ctx, id, Pack(event, data))
}

// BatchSendConn 批量发送多个连接
func (w *Server) BatchSendConn(ctx context.Context, ids []uint64, event string, data any) (err error) {
	msg := Pack(event, data)
	for _, id := range ids {
		if err = w.send(ctx, id, msg); err != nil {
			logger.Warnf("[ws.batchSend] err: %v", id, err)
		}
	}
	return err
}

// BatchSendMessage 批量发送多条消息
func (w *Server) BatchSendMessage(ctx context.Context, id uint64, list []string) error {
	if len(list) == 0 {
		return nil
	}
	conn, err := w.ws.GetManager(id).Get(id)
	if errors.Is(err, ws.ErrConnNotFound) {
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "[ws.Close] get conn by id: %v", id)
	}
	for _, msg := range list {
		if err = conn.AsyncSend(ctx, 0, []byte(msg)); err != nil {
			logger.Warnf("[ws.BatchSendMessage] AsyncSend connId:%d err: %v", id, err)
		}
	}

	return err
}

// Close 发送客户端连接
func (w *Server) Close(ctx context.Context, id uint64, data any) error {
	conn, err := w.ws.GetManager(id).Get(id)
	if errors.Is(err, ws.ErrConnNotFound) {
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "[ws.Close] get conn by id: %v", id)
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
		return w.send(ctx, cid, msg)
	})
	return
}

// send 发送消息至客户端
func (w *Server) send(ctx context.Context, id uint64, msg []byte) error {
	conn, err := w.ws.GetManager(id).Get(id)
	if errors.Is(err, ws.ErrConnNotFound) {
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "[ws.Close] get conn by id: %v", id)
	}

	if err = conn.AsyncSend(ctx, 0, msg); err != nil {
		return errors.Wrapf(err, "[ws.send] AsyncSend msg:%v", msg)
	}
	return nil
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

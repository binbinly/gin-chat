package service

import (
	"context"
	"encoding/json"
	"gin-chat/internal/websocket"
	"gin-chat/pkg/app"
	"time"

	"github.com/binbinly/pkg/auth"
	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"

	"gin-chat/pkg/redis"
)

// Online 用户在线服务接口
type Online interface {
	// UserKickOut 踢出上次登录
	UserKickOut(ctx context.Context, uid int) error
	// CheckOnline 检查用户是否在线
	CheckOnline(ctx context.Context, uid int) (bool, error)
	// UserOnline 用户上线
	UserOnline(ctx context.Context, token, sid string, cid uint64) (int, error)
	// UserOffline 用户下线
	UserOffline(ctx context.Context, uid int) error
	// GetUserConn 获取用户所在连接
	GetUserConn(ctx context.Context, uid int) *websocket.UserConnInfo
	// BatchUserConn 批量获取用户所在的连接
	BatchUserConn(ctx context.Context, uIds []int) ([]*websocket.UserConnInfo, error)
	// PushHistory 发送历史消息
	PushHistory(ctx context.Context, uid int) error
	// UserTokenCheck 用户token验证
	UserTokenCheck(ctx context.Context, token string) (int, error)
}

// UserKickOut 踢出上次登录
func (s *Service) UserKickOut(ctx context.Context, uid int) error {
	conn := s.GetUserConn(ctx, uid)
	if conn.ConnID == 0 {
		return nil
	}
	return s.ws.Close(ctx, conn, msgKickOut)
}

// UserOnline 用户上线
func (s *Service) UserOnline(ctx context.Context, token, sid string, cid uint64) (int, error) {
	uid, err := s.UserTokenCheck(ctx, token)
	if err != nil {
		return 0, err
	}
	//设置用户在线状态数据
	v, _ := json.Marshal(&websocket.UserConnInfo{
		UserID:   uid,
		ConnID:   cid,
		ServerID: sid,
	})
	if err = s.rdb.Set(ctx, app.BuildOnlineKey(uid), v, time.Duration(s.opts.jwtTimeout)*time.Second).Err(); err != nil {
		return 0, errors.Wrapf(err, "[service.online] user online hset err")
	}

	return uid, nil
}

// UserTokenCheck 用户token验证
func (s *Service) UserTokenCheck(ctx context.Context, token string) (int, error) {
	if len(token) == 0 {
		return 0, ErrUserTokenEmpty
	}
	p, err := auth.Parse(token, s.opts.jwtSecret)
	if err != nil {
		return 0, ErrUserTokenError
	}
	//获取当前合法用户token，检查此令牌是否已过期
	curToken := s.rdb.Get(ctx, app.BuildUserTokenKey(p.UserID)).Val()
	if curToken != token {
		return 0, ErrUserTokenExpired
	}

	return p.UserID, nil
}

// PushHistory 发送历史消息
func (s *Service) PushHistory(ctx context.Context, uid int) error {
	key := app.BuildHistoryKey(uid)
	list, err := s.rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return errors.Wrapf(err, "[service.online] lrange err: %v", err)
	}
	if err = s.ws.BatchSendMessage(ctx, s.GetUserConn(ctx, uid), list); err != nil {
		logger.Warnf("[service.online] send history msg err:%v", err)
	}

	return s.rdb.Del(ctx, key).Err()
}

// UserOffline 用户下线
func (s *Service) UserOffline(ctx context.Context, uid int) error {
	return s.rdb.Del(ctx, app.BuildOnlineKey(uid)).Err()
}

// CheckOnline 检查用户是否在线
func (s *Service) CheckOnline(ctx context.Context, uid int) (bool, error) {
	res, err := s.rdb.Exists(ctx, app.BuildOnlineKey(uid)).Result()
	if err != nil {
		return false, errors.Wrapf(err, "[service.online] check online id:%d", uid)
	}
	return res == redis.Success, nil
}

// GetUserConn 获取用户所在连接
func (s *Service) GetUserConn(ctx context.Context, uid int) *websocket.UserConnInfo {
	v := s.rdb.Get(ctx, app.BuildOnlineKey(uid)).Val()
	if v == "" {
		return &websocket.UserConnInfo{UserID: uid}
	}
	conn := &websocket.UserConnInfo{}
	_ = json.Unmarshal([]byte(v), conn)
	return conn
}

// BatchUserConn 批量获取用户所在的连接
func (s *Service) BatchUserConn(ctx context.Context, uIds []int) ([]*websocket.UserConnInfo, error) {
	if len(uIds) == 0 {
		return []*websocket.UserConnInfo{}, nil
	}
	keys := make([]string, len(uIds))
	for i, uid := range uIds {
		keys[i] = app.BuildOnlineKey(uid)
	}
	list, err := s.rdb.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[service.online[ get keys:%v by redis", keys)
	}
	conns := make([]*websocket.UserConnInfo, 0, len(list))
	for k, v := range list {
		switch val := v.(type) {
		case string:
			conn := &websocket.UserConnInfo{}
			_ = json.Unmarshal([]byte(val), conn)
			conns = append(conns, conn)
		default:
			conns = append(conns, &websocket.UserConnInfo{UserID: uIds[k]})
		}
	}
	return conns, nil
}

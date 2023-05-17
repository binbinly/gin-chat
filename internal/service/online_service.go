package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"

	"gin-chat/pkg/auth"
	"gin-chat/pkg/redis"
)

// UserConnInfo 用户连接信息
type UserConnInfo struct {
	ConnID   uint64 `json:"conn_id"`
	ServerID string `json:"server_id"`
}

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
	GetUserConn(ctx context.Context, uid int) *UserConnInfo
	// BatchUserConn 批量获取用户所在的连接
	BatchUserConn(ctx context.Context, uIds []int) ([]*UserConnInfo, error)
}

// UserKickOut 踢出上次登录
func (s *Service) UserKickOut(ctx context.Context, uid int) error {
	conn := s.GetUserConn(ctx, uid)
	if conn == nil {
		return nil
	}
	return s.ws.Close(ctx, conn.ConnID, msgKickOut)
}

// UserOnline 用户上线
func (s *Service) UserOnline(ctx context.Context, token, sid string, cid uint64) (int, error) {
	p, err := auth.Parse(token, s.opts.jwtSecret)
	if err != nil {
		return 0, ErrUserTokenError
	}
	//获取当前合法用户token，检查此令牌是否已过期
	curToken := s.rdb.Get(ctx, BuildUserTokenKey(p.UserID)).Val()
	if curToken != token {
		return 0, ErrUserTokenExpired
	}
	//设置用户在线状态数据
	v, _ := json.Marshal(&UserConnInfo{
		ConnID:   cid,
		ServerID: sid,
	})
	if err = s.rdb.Set(ctx, BuildOnlineKey(p.UserID), v, time.Duration(s.opts.jwtTimeout)*time.Second).Err(); err != nil {
		return 0, errors.Wrapf(err, "[service.online] user online hset err")
	}
	//获取当前用户的离线消息并发送
	list, err := s.rdb.LRange(ctx, BuildHistoryKey(p.UserID), 0, -1).Result()
	if err != nil {
		return 0, errors.Wrapf(err, "[service.online] lrange err: %v", err)
	}
	if err = s.ws.BatchSendMessage(ctx, cid, list); err != nil {
		logger.Warnf("[service.online] send history msg err:%v", err)
	}
	return p.UserID, nil
}

// UserOffline 用户下线
func (s *Service) UserOffline(ctx context.Context, uid int) error {
	return s.rdb.Del(ctx, BuildOnlineKey(uid)).Err()
}

// CheckOnline 检查用户是否在线
func (s *Service) CheckOnline(ctx context.Context, uid int) (bool, error) {
	res, err := s.rdb.Exists(ctx, BuildOnlineKey(uid)).Result()
	if err != nil {
		return false, errors.Wrapf(err, "[service.online] check online id:%d", uid)
	}
	return res == redis.Success, nil
}

// GetUserConn 获取用户所在连接
func (s *Service) GetUserConn(ctx context.Context, uid int) *UserConnInfo {
	v := s.rdb.Get(ctx, BuildOnlineKey(uid)).Val()
	if v == "" {
		return nil
	}
	conn := &UserConnInfo{}
	_ = json.Unmarshal([]byte(v), conn)
	return conn
}

// BatchUserConn 批量获取用户所在的连接
func (s *Service) BatchUserConn(ctx context.Context, uIds []int) ([]*UserConnInfo, error) {
	if len(uIds) == 0 {
		return []*UserConnInfo{}, nil
	}
	keys := make([]string, len(uIds))
	for i, uid := range uIds {
		keys[i] = BuildOnlineKey(uid)
	}
	list, err := s.rdb.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[service.online[ get keys:%v by redis", keys)
	}
	conns := make([]*UserConnInfo, 0, len(list))
	for _, v := range list {
		switch val := v.(type) {
		case string:
			conn := &UserConnInfo{}
			_ = json.Unmarshal([]byte(val), conn)
			conns = append(conns, conn)
		}
	}
	return conns, nil
}

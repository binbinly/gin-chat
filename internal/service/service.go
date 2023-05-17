package service

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"gin-chat/internal/model"
	"gin-chat/internal/repository"
	"gin-chat/internal/websocket"
	"gin-chat/pkg/cache"
	"gin-chat/pkg/mysql"
	redis2 "gin-chat/pkg/redis"
	"gin-chat/pkg/transport/ws"
)

const (
	msgFriendCreate = "你们已经是好友了，可以开始聊天啦"
	msgKickOut      = "账号已在其他地方登录了!"
)

const (
	// _onlinePrefix 在线key前缀
	_onlinePrefix = "user:online:"
	// _userPrefix 用户令牌标识 用于单点登录
	_userPrefix = "user:token:"
	// _historyPrefix 离线消息前缀
	_historyPrefix = "history:message:%d"
)

// 用于触发编译期的接口的合理性检查机制
var _ IService = (*Service)(nil)

// IService 服务接口定义
type IService interface {
	User
	Collect
	Moment
	Apply
	Friend
	Group
	Chat
	Online

	SendSMS(ctx context.Context, phone string) (string, error)
	CheckVCode(ctx context.Context, phone int64, vCode string) error
	EmoticonCat(ctx context.Context) (list []*model.Emoticon, err error)
	Emoticon(ctx context.Context, cat string) (list []*model.Emoticon, err error)
	ReportCreate(ctx context.Context, UserID, friendID int, cType int8, cat, content string) error
	Close() error
}

var Svc IService

// Service struct
type Service struct {
	opts options
	repo repository.IRepo
	rdb  *redis.Client
	ws   *websocket.Server
}

// New init service
func New(ws ws.Server, opts ...Option) (s *Service) {
	rdb := redis2.New()
	s = &Service{
		opts: newOptions(opts...),
		repo: repository.New(mysql.NewDB(), cache.NewCache(rdb)),
		rdb:  rdb,
		ws:   websocket.New(ws),
	}
	Svc = s
	return s
}

// Close service
func (s *Service) Close() error {
	return s.rdb.Close()
}

// BuildHistoryKey 历史消息键
func BuildHistoryKey(uid int) string {
	return fmt.Sprintf(_historyPrefix, uid)
}

// BuildOnlineKey 用户在线键
func BuildOnlineKey(uid int) string {
	return fmt.Sprintf("%s%d", _onlinePrefix, uid)
}

// BuildUserTokenKey 用户令牌键
func BuildUserTokenKey(uid int) string {
	return fmt.Sprintf("%s%d", _userPrefix, uid)
}

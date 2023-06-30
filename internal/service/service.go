package service

import (
	"context"

	"gin-chat/internal/model"
	"gin-chat/internal/repository"
	"gin-chat/internal/websocket"
	"gin-chat/pkg/mysql"
	redis2 "gin-chat/pkg/redis"

	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/transport/ws"
	"github.com/redis/go-redis/v9"
)

const (
	msgFriendCreate = "你们已经是好友了，可以开始聊天啦"
	msgKickOut      = "账号已在其他地方登录了!"
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
		repo: repository.New(mysql.NewDB(), cache.NewRedisCache(rdb)),
		rdb:  rdb,
		ws:   websocket.New(ws, rdb),
	}
	Svc = s
	return s
}

// Close service
func (s *Service) Close() error {
	return s.rdb.Close()
}

package repository

import (
	"context"
	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/repo"
	"gorm.io/gorm"

	"gin-chat/internal/model"
)

var _ IRepo = (*Repo)(nil)

// IRepo 数据仓库接口
type IRepo interface {
	User
	Apply
	Collect
	Friend
	Group
	GroupUser
	Moment
	MomentComment
	MomentTimeline
	MomentLike
	UserTag

	GetEmoticonCatAll(ctx context.Context) (list []*model.Emoticon, err error)
	GetEmoticonListByCat(ctx context.Context, cat string) (list []*model.Emoticon, err error)
	ReportCreate(ctx context.Context, report *model.ReportModel) (int, error)
	ReportExistPending(ctx context.Context, targetID int) (bool, error)
	CreateMessage(ctx context.Context, message model.MessageModel) (id int, err error)
	Close() error
}

// Repo dbs struct
type Repo struct {
	repo.Repo
}

// New new a Dao and return
func New(db *gorm.DB, c cache.Cache) IRepo {
	return &Repo{repo.Repo{
		DB:    db,
		Cache: c,
	}}
}

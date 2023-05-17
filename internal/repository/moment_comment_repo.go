package repository

import (
	"context"
	"fmt"

	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-chat/internal/model"
)

// MomentComment 朋友圈评论接口
type MomentComment interface {
	// CommentCreate 创建
	CommentCreate(ctx context.Context, model *model.MomentCommentModel) (id int, err error)
	// GetCommentsByMomentID 动态评论用户列表
	GetCommentsByMomentID(ctx context.Context, momentID int) ([]*model.MomentCommentModel, error)
	// GetCommentsByMomentIds 朋友圈多条动态的评论列表
	GetCommentsByMomentIds(ctx context.Context, mIds []int) (map[int][]*model.MomentCommentModel, error)
}

// CommentCreate 创建
func (r *Repo) CommentCreate(ctx context.Context, model *model.MomentCommentModel) (id int, err error) {
	if err = r.db.WithContext(ctx).Create(model).Error; err != nil {
		return 0, errors.Wrapf(err, "[repo.moment_comment] create")
	}
	r.delCache(ctx, commentCacheKey(model.MomentID))
	return model.ID, nil
}

// GetCommentsByMomentID 获取动态下所有评论
func (r *Repo) GetCommentsByMomentID(ctx context.Context, momentID int) (list []*model.MomentCommentModel, err error) {
	if err = r.queryCache(ctx, commentCacheKey(momentID), &list, func(data any) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Where("moment_id=?", momentID).
			Order("id asc").Find(data).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.moment_comment] query cache")
	}
	return
}

// GetCommentsByMomentIds 朋友圈动态下指定动态评论列表
func (r *Repo) GetCommentsByMomentIds(ctx context.Context, mIds []int) (mComments map[int][]*model.MomentCommentModel, err error) {
	mComments = make(map[int][]*model.MomentCommentModel, len(mIds))

	keys := make([]string, 0, len(mIds))
	for _, id := range mIds {
		keys = append(keys, commentCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*[]*model.MomentCommentModel)
	if err = r.cache.MultiGet(ctx, keys, cacheMap, func() any {
		return &[]*model.MomentCommentModel{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.moment_comment] multi get cache data err")
	}

	// 查询未命中
	for _, id := range mIds {
		comments, ok := cacheMap[commentCacheKey(id)]
		if !ok {
			cs, err := r.GetCommentsByMomentID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.moment_comment] get comment err: %v", err)
				continue
			}
			comments = &cs
		}
		if len(*comments) == 0 {
			continue
		}
		mComments[id] = *comments
	}
	return mComments, nil
}

func commentCacheKey(mid int) string {
	return fmt.Sprintf("moment:comment:%d", mid)
}

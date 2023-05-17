package repository

import (
	"context"

	"github.com/pkg/errors"

	"gin-chat/internal/model"
)

// Collect 收藏接口
type Collect interface {
	// CollectCreate 创建
	CollectCreate(ctx context.Context, collect *model.CollectModel) (id int, err error)
	// GetCollectsByUserID 获取用户的收藏列表
	GetCollectsByUserID(ctx context.Context, userID int, offset, limit int) (list []*model.CollectModel, err error)
	// CollectDelete 删除收藏
	CollectDelete(ctx context.Context, userID, id int) (err error)
}

// CollectCreate 创建收藏
func (r *Repo) CollectCreate(ctx context.Context, collect *model.CollectModel) (id int, err error) {
	if err = r.db.WithContext(ctx).Create(collect).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.collect] create collect")
	}

	return collect.ID, nil
}

// GetCollectsByUserID 获取用户收藏列表
func (r *Repo) GetCollectsByUserID(ctx context.Context, userID int, offset, limit int) (list []*model.CollectModel, err error) {
	// 从数据库中获取
	if err = r.db.WithContext(ctx).Scopes(model.OffsetPage(offset, limit)).Where("user_id = ? ", userID).
		Order(model.DefaultOrder).Find(&list).Error; err != nil {
		return nil, errors.Wrap(err, "[repo.collect] query db")
	}
	return
}

// CollectDelete 删除收藏
func (r *Repo) CollectDelete(ctx context.Context, userID, id int) (err error) {
	if err = r.db.WithContext(ctx).Where("user_id=?", userID).Delete(&model.CollectModel{}, id).Error; err != nil {
		return errors.Wrapf(err, "[repo.collect] destroy err")
	}

	return nil
}

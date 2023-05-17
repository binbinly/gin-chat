package service

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"gin-chat/internal/model"
	"gin-chat/internal/websocket"
	"gin-chat/pkg/mysql"
)

// Apply 申请好友接口
type Apply interface {
	// ApplyFriend 申请好友
	ApplyFriend(ctx context.Context, uid, fid int, nickname string, lookMe, lookHim int8) (err error)
	// ApplyMyList 我的申请列表
	ApplyMyList(ctx context.Context, uid int, offset, limit int) (list []*model.ApplyModel, users []*model.UserModel, err error)
	// ApplyPendingCount 待处理申请数
	ApplyPendingCount(ctx context.Context, uid int) (c int64, err error)
	// ApplyHandle 申请处理
	ApplyHandle(ctx context.Context, uid, fid int, nickname string, lookMe, lookHim int8) (err error)
}

// ApplyFriend 添加好友
func (s *Service) ApplyFriend(ctx context.Context, uid, fid int, nickname string, lookMe, lookHim int8) error {
	info, err := s.applyInfo(ctx, uid, fid)
	if err != nil {
		return err
	}
	if info.ID > 0 && info.Status == model.ApplyStatusPending { // 已存在
		return ErrApplyExisted
	}
	apply := model.ApplyModel{
		UID:      model.UID{UserID: uid},
		FriendID: fid,
		Nickname: nickname,
		LookMe:   lookMe,
		LookHim:  lookHim,
	}
	if _, err = s.repo.ApplyCreate(ctx, apply); err != nil {
		return errors.Wrapf(err, "[service.apply] create")
	}
	// 获取连接信息，发送消息
	c := s.GetUserConn(ctx, fid)
	if err = s.ws.Send(ctx, c.ConnID, websocket.EventNotify, &websocket.Notify{Type: "apply"}); err != nil {
		return errors.Wrapf(err, "[service.apply] ws send to client")
	}
	return nil
}

// ApplyMyList 用户申请列表
func (s *Service) ApplyMyList(ctx context.Context, uid int, offset, limit int) (list []*model.ApplyModel, users []*model.UserModel, err error) {
	list, err = s.repo.GetApplysByUserID(ctx, uid, offset, limit)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[service.apply] GetApplys id:%d", uid)
	}
	if len(list) == 0 {
		return
	}
	// 批量获取用户信息
	users, err = s.batchUserinfo(ctx, s.applyUserIds(list))
	if err != nil {
		return nil, nil, err
	}
	return
}

// ApplyPendingCount 待处理申请数量
func (s *Service) ApplyPendingCount(ctx context.Context, uid int) (c int64, err error) {
	return s.repo.ApplyPendingCount(ctx, uid)
}

// ApplyHandle 处理好友申请通过
func (s *Service) ApplyHandle(ctx context.Context, uid, fid int, nickname string, lookMe, lookHim int8) (err error) {
	info, err := s.applyInfo(ctx, fid, uid)
	if err != nil {
		return err
	}
	if info.ID == 0 || info.Status != model.ApplyStatusPending {
		return ErrApplyNotFound
	}
	// 开启事务
	tx := mysql.DB.Begin()
	// 我对好友的模型
	u := &model.FriendModel{
		UID:      model.UID{UserID: uid},
		FriendID: fid,
		Nickname: nickname,
		LookMe:   lookMe,
		LookHim:  lookHim,
	}
	// 好友对我的模型
	f := &model.FriendModel{
		UID:      model.UID{UserID: info.UserID},
		FriendID: uid,
		Nickname: info.Nickname,
		LookMe:   info.LookMe,
		LookHim:  info.LookHim,
	}
	if err = s.repo.FriendBatchCreate(ctx, tx, []*model.FriendModel{u, f}); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.apply] insert into friend to friend err")
	}
	// 修改申请状态
	if err = s.repo.ApplyUpdateStatus(ctx, tx, info.ID, info.FriendID); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.apply] update apply status err")
	}
	auth, err := s.userinfo(ctx, uid)
	if err != nil {
		tx.Rollback()
		return err
	}
	fAuth, err := s.userinfo(ctx, fid)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.apply] tx commit")
	}
	//推送消息 -> 好友
	cf := s.GetUserConn(ctx, fid)
	if err = s.ws.Send(ctx, cf.ConnID, websocket.EventChat, websocket.Chat{
		From: &websocket.Sender{
			ID:     uid,
			Name:   info.Nickname,
			Avatar: auth.Avatar,
		},
		ChatType: model.MessageChatTypeUser,
		Type:     model.MessageTypeSystem,
		Content:  msgFriendCreate,
		T:        time.Now().Unix(),
	}); err != nil {
		return errors.Wrapf(err, "[service.apply] ws send friend: %v", fid)
	}
	//推送消息 -> 自己
	cu := s.GetUserConn(ctx, uid)
	if err = s.ws.Send(ctx, cu.ConnID, websocket.EventChat, websocket.Chat{
		From: &websocket.Sender{
			ID:     fid,
			Name:   nickname,
			Avatar: fAuth.Avatar,
		},
		ChatType: model.MessageChatTypeUser,
		Type:     model.MessageTypeSystem,
		Content:  msgFriendCreate,
		T:        time.Now().Unix(),
	}); err != nil {
		return errors.Wrapf(err, "[service.apply] ws send self: %v", uid)
	}
	return nil
}

// applyInfo 申请详情
func (s *Service) applyInfo(ctx context.Context, uid, fid int) (apply *model.ApplyModel, err error) {
	apply, err = s.repo.GetApplyByFriendID(ctx, uid, fid)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.apply] find uid:%d,fid:%d", uid, fid)
	}
	return
}

// applyUserIds 申请人id列表
func (s *Service) applyUserIds(list []*model.ApplyModel) []int {
	ids := make([]int, 0, len(list))
	for _, apply := range list {
		ids = append(ids, apply.UserID)
	}
	return ids
}

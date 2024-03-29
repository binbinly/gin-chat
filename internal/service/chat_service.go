package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/xid"

	"gin-chat/internal/model"
	"gin-chat/internal/websocket"
	"gin-chat/pkg/app"
)

// Chat 聊天会话接口
type Chat interface {
	// ChatUserRecall 用户消息撤回
	ChatUserRecall(ctx context.Context, mid, id int, msgID string) (err error)
	// ChatGroupRecall 群消息撤回
	ChatGroupRecall(ctx context.Context, mid, id int, msgID string) (err error)
	// ChatUserDetail 好友聊天详情
	ChatUserDetail(ctx context.Context, mid, id int) (*websocket.Sender, error)
	// ChatGroupDetail 群聊天详情
	ChatGroupDetail(ctx context.Context, mid, id int) (*websocket.Sender, error)
	// ChatSendUser 发送单聊消息
	ChatSendUser(ctx context.Context, mid, id int, t int, content string, options json.RawMessage) (*websocket.Chat, error)
	// ChatSendGroup 发送群聊消息
	ChatSendGroup(ctx context.Context, mid, id int, t int, content string, options json.RawMessage) (*websocket.Chat, error)
}

// ChatUserRecall 用户消息撤回
func (s *Service) ChatUserRecall(ctx context.Context, mid, id int, msgID string) (err error) {
	_, err = s.friendInfo(ctx, id, mid)
	if err != nil {
		return err
	}
	// 发送消息
	if err = s.ws.Send(ctx, s.GetUserConn(ctx, id), websocket.EventRecall, &websocket.Recall{
		ID:       msgID,
		FromID:   mid,
		ToID:     id,
		ChatType: model.MessageChatTypeUser,
	}); err != nil {
		return errors.Wrapf(err, "[service.chat] send recall to user")
	}
	return nil
}

// ChatGroupRecall 群消息撤回
func (s *Service) ChatGroupRecall(ctx context.Context, mid, id int, msgID string) (err error) {
	if !s.checkOnline(ctx, mid) {
		return ErrUserOffline
	}
	users, err := s.repo.GroupUserAll(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.chat] group user all id:%d", id)
	}
	msg := &websocket.Recall{
		ID:       msgID,
		FromID:   mid,
		ToID:     id,
		ChatType: model.MessageChatTypeGroup,
	}
	// 没有群成员，说明是聊天室，直接广播在线用户
	if users == nil || len(users) == 0 {
		if err = s.ws.Broadcast(ctx, mid, websocket.EventRecall, msg); err != nil {
			return errors.Wrapf(err, "[service.chat] ws broadcast group")
		}
		return nil
	}

	ids := make([]int, 0, len(users))
	for _, u := range users {
		if u.UserID == mid { // 不需要推送自己
			continue
		}
		ids = append(ids, u.UserID)
	}
	if len(ids) == len(users) { //群中没有自己
		return ErrGroupUserNotJoin
	}
	cs, err := s.BatchUserConn(ctx, ids)
	if err != nil {
		return err
	}
	// 发送消息
	if err = s.ws.BatchSendConn(ctx, cs, websocket.EventRecall, msg); err != nil {
		return errors.Wrapf(err, "[service.group] ws send recall to group")
	}
	return nil
}

// ChatUserDetail 好友聊天详情
func (s *Service) ChatUserDetail(ctx context.Context, mid, id int) (*websocket.Sender, error) {
	// 好友->我关系详情,判断是否为好友
	f, err := s.friendInfo(ctx, id, mid)
	if err != nil {
		return nil, err
	}
	// 判断对方是否拉黑你
	if f.IsBlack == 1 {
		return nil, ErrFriendNotFound
	}
	// 获取我备注的好友昵称
	m, err := s.friendInfo(ctx, mid, id)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	// 好友用户详情
	u, err := s.userinfo(ctx, id)
	if err != nil {
		return nil, err
	}
	return &websocket.Sender{
		ID:     u.ID,
		Name:   m.Nickname,
		Avatar: app.BuildResUrl(u.Avatar),
	}, nil
}

// ChatGroupDetail 群聊天详情
func (s *Service) ChatGroupDetail(ctx context.Context, mid, id int) (*websocket.Sender, error) {
	group, err := s.groupInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	if group.Type == model.GroupTypeGroup {
		//是否是群成员
		if err = s.isGroupUser(ctx, mid, id); err != nil {
			return nil, err
		}
	}

	return &websocket.Sender{
		ID:     group.ID,
		Name:   group.Name,
		Avatar: group.Avatar,
	}, nil
}

// ChatSendUser 发送单聊消息
func (s *Service) ChatSendUser(ctx context.Context, mid, id int, t int, content string, options json.RawMessage) (*websocket.Chat, error) {
	if !s.checkOnline(ctx, mid) {
		return nil, ErrUserOffline
	}
	f, err := s.friendInfo(ctx, id, mid)
	if err != nil {
		return nil, err
	}
	// 判断对方是否拉黑你
	if f.IsBlack == 1 {
		return nil, ErrFriendNotFound
	}
	// 我的用户详情
	u, err := s.userinfo(ctx, mid)
	if err != nil {
		return nil, err
	}
	//构建消息
	m := &websocket.Chat{
		ID: xid.New().String(),
		From: &websocket.Sender{
			ID:     u.ID,
			Name:   f.Nickname,
			Avatar: app.BuildResUrl(u.Avatar),
		},
		ChatType: model.MessageChatTypeUser,
		Type:     t,
		Content:  content,
		Options:  options,
		T:        time.Now().Unix(),
	}
	// 发送消息
	if err = s.ws.Send(ctx, s.GetUserConn(ctx, id), websocket.EventChat, m); err != nil {
		return nil, errors.Wrapf(err, "[service.chat] ws send to user")
	}
	return m, nil
}

// ChatSendGroup 发送群聊消息
func (s *Service) ChatSendGroup(ctx context.Context, mid, id int, t int, content string, options json.RawMessage) (*websocket.Chat, error) {
	if !s.checkOnline(ctx, mid) {
		return nil, ErrUserOffline
	}
	group, users, _, err := s.groupUsers(ctx, mid, id)
	if err != nil {
		return nil, err
	}
	u, err := s.userinfo(ctx, mid)
	if err != nil {
		return nil, err
	}
	//构建消息
	m := &websocket.Chat{
		ID: xid.New().String(),
		From: &websocket.Sender{
			ID:     u.ID,
			Name:   getNickname(u),
			Avatar: app.BuildResUrl(u.Avatar),
		},
		To: &websocket.Sender{
			ID:     group.ID,
			Name:   group.Name,
			Avatar: group.Avatar,
		},
		ChatType: model.MessageChatTypeGroup,
		Type:     t,
		Content:  content,
		Options:  options,
		T:        time.Now().Unix(),
	}

	// 没有群成员，说明是聊天室，直接广播在线用户
	if users == nil {
		if err = s.ws.Broadcast(ctx, mid, websocket.EventChat, m); err != nil {
			return nil, errors.Wrapf(err, "[service.chat] ws broadcast group")
		}
		return m, nil
	}
	userIds := make([]int, 0, len(users))
	for _, user := range users {
		if user.UserID == mid { // 当前用户消息返回，不用ws推送
			if user.Nickname != "" { //设置了群昵称
				m.From.Name = user.Nickname
			}
			continue
		}
		userIds = append(userIds, user.UserID)
	}
	// 获取连接
	cs, err := s.BatchUserConn(ctx, userIds)
	if err != nil {
		return nil, err
	}
	// 推送消息
	if err = s.ws.BatchSendConn(ctx, cs, websocket.EventChat, m); err != nil {
		return nil, errors.Wrapf(err, "[service.chat] ws send to group")
	}
	return m, nil
}

// checkOnline 检查用户是否在线
func (s *Service) checkOnline(ctx context.Context, uid int) bool {
	c := s.GetUserConn(ctx, uid)
	if c == nil {
		return false
	}
	return true
}

func getNickname(u *model.UserModel) string {
	name := u.Username
	if u.Nickname != "" {
		name = u.Nickname
	}
	return name
}

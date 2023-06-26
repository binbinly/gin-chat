package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"

	"gin-chat/internal/model"
	"gin-chat/internal/websocket"
	"gin-chat/pkg/mysql"
)

const (
	msgCreate     = "群聊已创建，可以开始聊天啦"
	msgEditName   = "修改群聊名为 %s"
	msgEditRemark = "[新公告] %s"
	msgKickoff    = "将 %s 移出了群聊"
	msgInvite     = "邀请 %s 加入了群聊"
	msgJoin       = "加入了群聊"
	msgQuit       = "退出了该群聊"
	msgDisband    = "解散了群聊"
)

// 发送消息结构体
type sendParams struct {
	userID   int                     // 操作人ID
	group    *model.GroupModel       // 群模型
	gUsers   []*model.GroupUserModel // 群成员模型数组
	content  string                  // 推送消息内容
	targetID int                     // 目标人
	tContent string                  // 目标人消息
}

// GroupDetail 群详情结构
type GroupDetail struct {
	Group      *model.GroupModel       //群信息
	GroupUsers []*model.GroupUserModel //群成员列表
	Users      []*model.UserModel      //群成员用户信息列表
	My         *model.GroupUserModel   //我的群成员信息
}

// Group 群组服务接口
type Group interface {
	// GroupCreate 创建群组
	GroupCreate(ctx context.Context, mid int, ids []int) error
	// GroupEditName 修改群组名
	GroupEditName(ctx context.Context, mid, id int, name string) error
	// GroupEditRemark 修噶群公告
	GroupEditRemark(ctx context.Context, mid, id int, remark string) error
	// GroupEditUserNickname 修改我的群昵称
	GroupEditUserNickname(ctx context.Context, mid, id int, nickname string) error
	// GroupMyList 我的群列表
	GroupMyList(ctx context.Context, mid int) ([]*model.GroupList, error)
	// GroupInfo 群详情
	GroupInfo(ctx context.Context, mid, id int) (*GroupDetail, error)
	// GroupUserAll 群成员
	GroupUserAll(ctx context.Context, mid, id int) ([]*model.GroupUserModel, []*model.UserModel, error)
	// GroupUserQuit 退出群
	GroupUserQuit(ctx context.Context, mid, id int) error
	// GroupKickOffUser 踢出群
	GroupKickOffUser(ctx context.Context, mid, id, tid int) error
	// GroupInviteUser 邀请入群
	GroupInviteUser(ctx context.Context, mid, id, tid int) error
	// GroupJoin 加入群
	GroupJoin(ctx context.Context, mid, id int) error
}

// GroupCreate 创建群组
func (s *Service) GroupCreate(ctx context.Context, mid int, ids []int) error {
	u, err := s.userinfo(ctx, mid)
	if err != nil {
		return err
	}
	uname := getNickname(u)
	// 选择好友信息
	friends, err := s.friendsByIds(ctx, mid, ids)
	if err != nil {
		return err
	}
	if len(friends) == 0 {
		return ErrFriendNotRecord
	}
	// 批量获取好友信息
	fids := friendIds(friends)
	users, err := s.batchUserinfo(ctx, fids)
	if err != nil {
		return err
	}
	group := &model.GroupModel{
		UID:  model.UID{UserID: mid},
		Name: buildGroupName(uname, users),
	}
	// 开启事务
	tx := mysql.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 创建群组
	gid, err := s.repo.GroupCreate(ctx, tx, group)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.group] create")
	}
	// 创建群组成员
	if err = s.repo.GroupUserBatchCreate(ctx, tx, buildGroupUsers(mid, gid, users)); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.group] users create")
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.group] tx commit err")
	}

	// 获取创建者，所有好友的连接
	fids = append(fids, mid)
	cs, err := s.BatchUserConn(ctx, fids)
	if err != nil {
		return err
	}
	// 发送消息
	if err = s.ws.BatchSendConn(ctx, cs, websocket.EventChat, &websocket.Chat{
		From: &websocket.Sender{
			ID:     u.ID,
			Name:   uname,
			Avatar: u.Avatar,
		},
		To: &websocket.Sender{
			ID:     group.ID,
			Name:   group.Name,
			Avatar: group.Avatar,
		},
		ChatType: model.MessageChatTypeGroup,
		Type:     model.MessageTypeSystem,
		Content:  msgCreate,
		T:        time.Now().Unix(),
	}); err != nil {
		return errors.Wrapf(err, "[service.group] ws send create")
	}
	return nil
}

// GroupEditName 更新群名
func (s *Service) GroupEditName(ctx context.Context, mid, id int, name string) error {
	group, gUsers, _, err := s.groupUsers(ctx, mid, id)
	if err != nil {
		return err
	}
	if group.Name == name { // 群名没变，无需修改
		return nil
	}
	group.Name = name
	// 修改群组信息
	if err = s.repo.GroupSave(ctx, group); err != nil {
		return errors.Wrapf(err, "[service.group] save by id: %d", id)
	}
	return s.sendMessage(ctx, &sendParams{
		userID:  mid,
		group:   group,
		gUsers:  gUsers,
		content: fmt.Sprintf(msgEditName, name),
	})
}

// GroupEditRemark 更新群公告
func (s *Service) GroupEditRemark(ctx context.Context, mid, id int, remark string) error {
	group, gUsers, _, err := s.groupUsers(ctx, mid, id)
	if err != nil {
		return err
	}
	if group.Remark == remark { // 公告没变，无需修改
		return ErrGroupDataUnmodified
	}
	group.Remark = remark
	// 修改群组信息
	if err = s.repo.GroupSave(ctx, group); err != nil {
		return errors.Wrapf(err, "[service.group] save group by id: %d", id)
	}
	return s.sendMessage(ctx, &sendParams{
		userID:  mid,
		group:   group,
		gUsers:  gUsers,
		content: fmt.Sprintf(msgEditRemark, remark),
	})
}

// GroupEditUserNickname 更新我在群组中的昵称
func (s *Service) GroupEditUserNickname(ctx context.Context, mid, id int, nickname string) error {
	if err := s.isGroupUser(ctx, mid, id); err != nil {
		return err
	}
	return s.repo.GroupUserUpdateNickname(ctx, mid, id, nickname)
}

// GroupMyList 我的群组
func (s *Service) GroupMyList(ctx context.Context, mid int) (list []*model.GroupList, err error) {
	return s.repo.GetGroupsByUserID(ctx, mid)
}

// GroupInfo 群组信息
func (s *Service) GroupInfo(ctx context.Context, mid, id int) (*GroupDetail, error) {
	group, gUsers, my, err := s.groupUsers(ctx, mid, id)
	if err != nil {
		return nil, err
	}
	// 批量获取用户信息
	users, err := s.batchUserinfo(ctx, s.groupUserIds(gUsers))
	if err != nil {
		return nil, err
	}
	return &GroupDetail{
		Group:      group,
		GroupUsers: gUsers,
		Users:      users,
		My:         my,
	}, nil
}

// GroupUserAll 所有群成员
func (s *Service) GroupUserAll(ctx context.Context, mid, id int) ([]*model.GroupUserModel, []*model.UserModel, error) {
	_, gUsers, _, err := s.groupUsers(ctx, mid, id)
	if err != nil {
		return nil, nil, err
	}
	// 批量获取用户信息
	users, err := s.batchUserinfo(ctx, s.groupUserIds(gUsers))
	if err != nil {
		return nil, nil, err
	}
	return gUsers, users, nil
}

// GroupUserQuit 删除并退出群聊
func (s *Service) GroupUserQuit(ctx context.Context, mid, id int) (err error) {
	group, gUsers, _, err := s.groupUsers(ctx, mid, id)
	if err != nil {
		return err
	}
	if group.UserID == mid { // 管理员退出群，直接解散群
		return s.deleteGroup(ctx, mid, group, gUsers)
	}
	// 删除群成员
	return s.deleteGroupUser(ctx, mid, group, gUsers)
}

// GroupKickOffUser 踢出群成员
func (s *Service) GroupKickOffUser(ctx context.Context, mid, id, tid int) error {
	group, gUsers, _, err := s.groupUsers(ctx, mid, id)
	if err != nil {
		return err
	}
	// 目标用户信息
	toUser, err := s.userinfo(ctx, tid)
	if err != nil {
		return err
	}
	// 目标用户是否为群成员
	to := inGroup(gUsers, tid)
	if to == nil {
		return ErrGroupUserTargetNotJoin
	}
	// 删除被踢成员
	if err = s.repo.GroupUserDelete(ctx, to); err != nil {
		return errors.Wrapf(err, "[service.group] kickoff err uid:%d,gid:%d", tid, id)
	}
	// 被踢人昵称
	kName := to.Nickname
	if kName == "" {
		if toUser.Nickname != "" {
			kName = toUser.Nickname
		} else {
			kName = toUser.Username
		}
	}
	// 发送消息
	return s.sendMessage(ctx, &sendParams{
		userID:   mid,
		group:    group,
		gUsers:   gUsers,
		content:  fmt.Sprintf(msgKickoff, kName),
		targetID: tid,
		tContent: fmt.Sprintf(msgKickoff, "你"),
	})
}

// GroupInviteUser 邀请好友加入
func (s *Service) GroupInviteUser(ctx context.Context, mid, id, tid int) (err error) {
	group, gUsers, _, err := s.groupUsers(ctx, mid, id)
	if err != nil {
		return err
	}
	// 目标用户信息
	toUser, err := s.userinfo(ctx, tid)
	if err != nil {
		return err
	}
	// 目标用户是否已经为群成员
	to := inGroup(gUsers, tid)
	if to != nil {
		return ErrGroupUserExisted
	}
	// 加入群聊
	gUser := &model.GroupUserModel{
		UID:     model.UID{UserID: tid},
		GroupID: group.ID,
	}

	if err = s.repo.GroupUserCreate(ctx, gUser); err != nil {
		return errors.Wrapf(err, "[service.group] CreateUser uid:%d, gid:%d", tid, id)
	}
	name := toUser.Username
	if toUser.Nickname != "" {
		name = toUser.Nickname
	}
	gUsers = append(gUsers, gUser)
	return s.sendMessage(ctx, &sendParams{
		userID:   mid,
		group:    group,
		gUsers:   gUsers,
		content:  fmt.Sprintf(msgInvite, name),
		targetID: tid,
		tContent: fmt.Sprintf(msgInvite, "你"),
	})
}

// GroupJoin 加入群
func (s *Service) GroupJoin(ctx context.Context, uid, id int) error {
	group, gUsers, _, err := s.groupUsers(ctx, uid, id)
	if err != nil {
		return err
	}
	// 加入群聊
	gUser := &model.GroupUserModel{
		UID:     model.UID{UserID: uid},
		GroupID: group.ID,
	}

	if err = s.repo.GroupUserCreate(ctx, gUser); err != nil {
		return errors.Wrapf(err, "[service.group] CreateUser uid:%d, gid:%d", uid, id)
	}
	gUsers = append(gUsers, gUser)
	return s.sendMessage(ctx, &sendParams{
		userID:  uid,
		group:   group,
		gUsers:  gUsers,
		content: msgJoin,
	})
}

// sendMessage 发送群消息
func (s *Service) sendMessage(ctx context.Context, params *sendParams) (err error) {
	mContent := params.content
	// 我的用户详情
	m, err := s.userinfo(ctx, params.userID)
	if err != nil {
		return err
	}
	params.content = fmt.Sprintf("%s %s", s.myGroupName(m, params.gUsers), params.content)

	f := &websocket.Sender{
		ID:     m.ID,
		Name:   s.myGroupName(m, params.gUsers),
		Avatar: m.Avatar,
	}
	t := &websocket.Sender{
		ID:     params.group.ID,
		Name:   params.group.Name,
		Avatar: params.group.Avatar,
	}
	now := time.Now().Unix()
	// 给群组成员发送消息
	for _, gUser := range params.gUsers {
		ct := params.content
		if gUser.UserID == params.userID { // 发送给自己的消息
			ct = "你 " + mContent
		}
		if gUser.UserID == params.targetID {
			ct = fmt.Sprintf("%s %s", s.myGroupName(m, params.gUsers), params.tContent)
		}

		if err = s.ws.Send(ctx, s.GetUserConn(ctx, gUser.UserID), websocket.EventChat, &websocket.Chat{
			From:     f,
			To:       t,
			ChatType: model.MessageChatTypeGroup,
			Type:     model.MessageTypeSystem,
			Content:  ct,
			T:        now,
		}); err != nil {
			logger.Warnf("[service.group] ws send uid: %v, err: %v", gUser.UserID, err)
		}
	}
	return nil
}

// myGroupName 我在群组中显示的昵称
func (s *Service) myGroupName(my *model.UserModel, users []*model.GroupUserModel) string {
	name := my.Username
	if my.Nickname != "" {
		name = my.Nickname
	}
	// 群组信息中是否单独设置了昵称
	for _, u := range users {
		if u.UserID == my.ID {
			if u.Nickname != "" {
				name = u.Nickname
			}
			break
		}
	}
	return name
}

// groupUsers 获取群，群成员信息
func (s *Service) groupUsers(ctx context.Context, uid, gid int) (*model.GroupModel, []*model.GroupUserModel, *model.GroupUserModel, error) {
	// 群信息
	group, err := s.groupInfo(ctx, gid)
	if err != nil {
		return nil, nil, nil, err
	}
	// 群组成员
	gUsers, err := s.repo.GroupUserAll(ctx, gid)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "[service.group] user all id:%d", gid)
	}
	my := inGroup(gUsers, uid)
	if my == nil {
		return nil, nil, nil, ErrGroupUserNotJoin
	}
	return group, gUsers, my, nil
}

// groupInfo 获取群信息
func (s *Service) groupInfo(ctx context.Context, id int) (group *model.GroupModel, err error) {
	group, err = s.repo.GetGroupByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.group] info id:%d", id)
	}
	if group.ID == 0 {
		return nil, ErrGroupNotFound
	}
	return
}

// isGroupUser 是否是群成员
func (s *Service) isGroupUser(ctx context.Context, uid, gid int) error {
	users, err := s.repo.GroupUserAll(ctx, gid)
	if err != nil {
		return errors.Wrapf(err, "[service.chat] group user all: %d", gid)
	}
	if u := inGroup(users, uid); u != nil {
		return nil
	}
	return ErrGroupUserNotJoin
}

// deleteGroupUser 删除群成员
func (s *Service) deleteGroupUser(ctx context.Context, uid int, group *model.GroupModel, gUsers []*model.GroupUserModel) (err error) {
	t := inGroup(gUsers, uid)
	if err = s.repo.GroupUserDelete(ctx, t); err != nil {
		return errors.Wrapf(err, "[service.group] quit err uid:%d,gid:%d", uid, t.GroupID)
	}
	return s.sendMessage(ctx, &sendParams{
		userID:  uid,
		group:   group,
		gUsers:  gUsers,
		content: msgQuit,
	})
}

// deleteGroup 删除群组
func (s *Service) deleteGroup(ctx context.Context, uid int, group *model.GroupModel, gUsers []*model.GroupUserModel) (err error) {
	// 开启事务
	tx := mysql.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 删除群
	if err = s.repo.GroupDelete(ctx, tx, group); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.group] delete err")
	}
	// 删除群成员
	if err = s.repo.GroupUserDeleteByGroupID(ctx, tx, group.ID); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.group] delete users err")
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "tx commit err")
	}
	// 发送消息
	return s.sendMessage(ctx, &sendParams{
		userID:  uid,
		group:   group,
		gUsers:  gUsers,
		content: msgDisband,
	})
}

// groupUserIds 获取群成员所有id
func (s *Service) groupUserIds(users []*model.GroupUserModel) (ids []int) {
	ids = make([]int, 0, len(users))
	for _, u := range users {
		ids = append(ids, u.UserID)
	}
	return
}

// buildGroupUsers 构建批量创建群成员结构
func buildGroupUsers(uid, id int, friends []*model.UserModel) []*model.GroupUserModel {
	users := make([]*model.GroupUserModel, 0)
	// 自己加入群组
	users = append(users, &model.GroupUserModel{
		UID:     model.UID{UserID: uid},
		GroupID: id,
	})
	for _, friend := range friends {
		users = append(users, &model.GroupUserModel{
			UID:     model.UID{UserID: friend.ID},
			GroupID: id,
		})
	}
	return users
}

// inGroup 是否是群组内成员，并返回当前群成员信息
func inGroup(users []*model.GroupUserModel, uid int) *model.GroupUserModel {
	for _, user := range users {
		if user.UserID == uid {
			return user
		}
	}
	return nil
}

// buildGroupName 构建群组默认名称
func buildGroupName(username string, friends []*model.UserModel) string {
	var name strings.Builder
	name.WriteString(username)
	for i, f := range friends {
		if i == 4 { //最多拼接4位好友昵称
			break
		}
		m := f.Username
		if f.Nickname != "" {
			m = f.Nickname
		}
		name.WriteString(",")
		name.WriteString(m)
	}
	return name.String()
}

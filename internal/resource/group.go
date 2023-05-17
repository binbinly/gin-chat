package resource

import (
	"gin-chat/internal/model"
)

// GroupResponse 群组响应结构
type GroupResponse struct {
	Nickname string        `json:"nickname"`
	Info     *Group        `json:"info"`
	Users    []*model.User `json:"users"`
}

// Group 群详情结构
type Group struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	InviteConfirm int8   `json:"invite_confirm"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	Remark        string `json:"remark"`
}

// GroupResource 群组信息转换
func GroupResource(group *model.GroupModel, users []*model.UserModel, gUsers []*model.GroupUserModel,
	my *model.GroupUserModel) *GroupResponse {
	return &GroupResponse{
		Info: &Group{
			ID:            group.ID,
			UserID:        group.UserID,
			InviteConfirm: group.InviteConfirm,
			Name:          group.Name,
			Avatar:        group.Avatar,
			Remark:        group.Remark,
		},
		Nickname: my.Nickname,
		Users:    GroupUsersResource(users[:4], gUsers), // 只需要返回4个用户就可以了
	}
}

// GroupUsersResource 群成员列表转换
func GroupUsersResource(users []*model.UserModel, gUsers []*model.GroupUserModel) []*model.User {
	us := make([]*model.User, 0)
	uMap := gUserToMap(gUsers)
	for _, user := range users {
		name := user.Username
		if user.Nickname != "" { // 设置了昵称
			name = user.Nickname
		}
		if nick, ok := uMap[user.ID]; ok {
			name = nick
		}
		us = append(us, &model.User{
			ID:     user.ID,
			Name:   name,
			Avatar: user.Avatar,
		})
	}
	return us
}

// gUserToMap 转换群成员map结构 user_id => nickname
func gUserToMap(gUsers []*model.GroupUserModel) (m map[int]string) {
	m = make(map[int]string, len(gUsers))
	for _, user := range gUsers {
		if user.Nickname != "" {
			m[user.UserID] = user.Nickname
		}
	}
	return
}

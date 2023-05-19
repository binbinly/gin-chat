package resource

import "gin-chat/internal/model"

// FriendResponse 好友响应结构
type FriendResponse struct {
	Friend   *Friend       `json:"friend"`
	User     *UserResponse `json:"user"`
	IsFriend bool          `json:"is_friend"`
}

// Friend 好友结构
type Friend struct {
	LookMe   int8     `json:"look_me"`
	LookHim  int8     `json:"look_him"`
	IsStar   int8     `json:"is_star"`
	IsBlack  int8     `json:"is_black"`
	Nickname string   `json:"nickname"`
	Tags     []string `json:"tags"`
}

// FriendResource 好友结构转换
func FriendResource(f *model.FriendModel, u *model.UserModel, tags []string) *FriendResponse {
	info := &FriendResponse{
		User: UserResource(u),
	}
	if f == nil || f.ID == 0 {
		info.Friend = nil
	} else {
		info.IsFriend = true
		info.Friend = &Friend{
			Nickname: f.Nickname,
			LookMe:   f.LookMe,
			LookHim:  f.LookHim,
			IsStar:   f.IsStar,
			IsBlack:  f.IsBlack,
			Tags:     tags,
		}
	}
	return info
}

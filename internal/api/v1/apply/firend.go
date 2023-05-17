package apply

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/ecode"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// friendParams 申请好友
type friendParams struct {
	FriendID int    `json:"friend_id" binding:"required,numeric" example:"1"`         //好友ID
	Nickname string `json:"nickname"  binding:"required,min=1,max=30" example:"test"` //备注昵称
	LookMe   int8   `json:"look_me"  binding:"required,oneof=0 1" example:"1"`        //看我
	LookHim  int8   `json:"look_him" binding:"required,oneof=0 1" example:"1"`        //看他
}

// Friend 申请好友
// @Summary 申请好友
// @Description 申请好友
// @Tags 好友申请
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body friendParams true "friend"
// @success 0 {object} app.Response "调用成功结构"
// @Router /apply/friend [post]
func Friend(c *gin.Context) {
	var req friendParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	uid := api.GetUserID(c)
	if uid == req.FriendID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.ApplyFriend(c.Request.Context(), uid, req.FriendID, req.Nickname, req.LookMe, req.LookHim)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

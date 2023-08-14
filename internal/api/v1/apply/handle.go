package apply

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/ecode"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// handleParams 处理好友申请
type handleParams struct {
	FriendID int    `json:"friend_id" binding:"required,numeric" example:"1"`         //好友ID
	Nickname string `json:"nickname"  binding:"required,min=1,max=30" example:"test"` //备注内侧
	LookMe   int8   `json:"look_me"  binding:"required,oneof=0 1" example:"1"`        //看我
	LookHim  int8   `json:"look_him" binding:"required,oneof=0 1" example:"1"`        //看他
}

// Handle 处理好友申请
// @Summary 处理好友申请
// @Description 处理好友申请
// @Tags 好友申请
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body handleParams true "handle"
// @success 0 {object} app.Response "调用成功结构"
// @Router /apply/handle [post]
func Handle(c *gin.Context) {
	var req handleParams
	if err := api.BindJSON(c, &req); err != nil {
		app.ErrorParamInvalid(c, err)
		return
	}

	uid := api.GetUserID(c)
	if uid == req.FriendID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.ApplyHandle(c.Request.Context(), uid, req.FriendID, req.Nickname, req.LookMe, req.LookHim)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

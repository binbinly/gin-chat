package group

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// actionParams 操作群用户
type actionParams struct {
	ID     int `json:"id" binding:"required,numeric" example:"1"`      // 群ID
	UserID int `json:"user_id" binding:"required,numeric" example:"1"` // 用户ID
}

// Invite 邀请好友
// @Summary 邀请好友
// @Description 邀请好友
// @Tags 群组
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body actionParams true "The group info"
// @success 0 {object} app.Response "调用成功结构"
// @Router /group/invite [post]
func Invite(c *gin.Context) {
	var req actionParams
	if err := api.BindJSON(c, &req); err != nil {
		app.ErrorParamInvalid(c, err)
		return
	}

	err := service.Svc.GroupInviteUser(c.Request.Context(), api.GetUserID(c), req.ID, req.UserID)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

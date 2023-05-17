package friend

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/ecode"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// destroyParams 删除好友
type destroyParams struct {
	UserID int `json:"user_id" binding:"required,numeric" example:"1"` // 用户ID
}

// Destroy 删除好友
// @Summary 删除好友
// @Description 删除好友
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body destroyParams true "destroy"
// @success 0 {object} app.Response "调用成功结构"
// @Router /friend/destroy [post]
func Destroy(c *gin.Context) {
	var req destroyParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	uid := api.GetUserID(c)
	if uid == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.FriendDestroy(c.Request.Context(), uid, req.UserID)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

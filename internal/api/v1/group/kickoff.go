package group

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// KickOff 踢出群成员
// @Summary 踢出群成员
// @Description 踢出群成员
// @Tags 群组
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body actionParams true "The group info"
// @success 0 {object} app.Response "调用成功结构"
// @Router /group/kickoff [post]
func KickOff(c *gin.Context) {
	var req actionParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.GroupKickOffUser(c.Request.Context(), api.GetUserID(c), req.ID, req.UserID)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

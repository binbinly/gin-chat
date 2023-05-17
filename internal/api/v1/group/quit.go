package group

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// Quit 删除并退出群
// @Summary 删除并退出群
// @Description 删除并退出群
// @Tags 群组
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id query int true "群ID"
// @success 0 {object} app.Response "调用成功结构"
// @Router /group/quit [get]
func Quit(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	if id == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.GroupUserQuit(c.Request.Context(), api.GetUserID(c), id)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

package group

import (
	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/util"
	"github.com/gin-gonic/gin"
)

// Join 加入群
// @Summary 加入群
// @Description 加入群
// @Tags 群组
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id query int true "群ID"
// @success 0 {object} app.Response "调用成功结构"
// @Router /group/join [get]
func Join(c *gin.Context) {
	id := util.MustInt(c.Query("id"))
	if id == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.GroupJoin(c.Request.Context(), api.GetUserID(c), id)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

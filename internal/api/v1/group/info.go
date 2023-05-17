package group

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"gin-chat/internal/api"
	"gin-chat/internal/resource"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// Info 获取群信息
// @Summary 获取群信息
// @Description 获取群信息
// @Tags 群组
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id query int true "群ID"
// @success 0 {object} app.Response{data=resource.GroupResponse} "调用成功结构"
// @Router /group/info [get]
func Info(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	if id == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	info, err := service.Svc.GroupInfo(c.Request.Context(), api.GetUserID(c), id)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, resource.GroupResource(info.Group, info.Users, info.GroupUsers, info.My))
}

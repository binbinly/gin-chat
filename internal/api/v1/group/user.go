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

// User 获取群成员
// @Summary 获取群成员
// @Description 获取群成员
// @Tags 群组
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id query int true "群ID"
// @success 0 {object} app.Response{data=[]model.User} "调用成功结构"
// @Router /group/user [get]
func User(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	if id == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	gUsers, users, err := service.Svc.GroupUserAll(c.Request.Context(), api.GetUserID(c), id)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, resource.GroupUsersResource(users, gUsers, 0))
}

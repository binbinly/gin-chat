package user

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/resource"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// Profile 获取用户信息
// @Summary 获取个人资料
// @Description 获取个人资料
// @Tags 用户
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=resource.UserResponse} "调用成功结构"
// @Router /user/profile [get]
func Profile(c *gin.Context) {
	user, err := service.Svc.UserInfoByID(c.Request.Context(), api.GetUserID(c))
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, resource.UserResource(user))
}

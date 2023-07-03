package moment

import (
	"gin-chat/internal/api"
	"gin-chat/internal/resource"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"

	"github.com/binbinly/pkg/util"
	"github.com/gin-gonic/gin"
)

// List 动态列表
// @Summary 动态列表
// @Description 动态列表
// @Tags 朋友圈
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param user_id query int false "用户id"
// @Param p query int false "页码"
// @success 0 {object} app.Response{data=[]model.Moment} "调用成功结构"
// @Router /moment/list [get]
func List(c *gin.Context) {
	mid := api.GetUserID(c)
	uid := util.MustInt(c.Query("user_id"))
	if uid == 0 { // 默认查看自己的动态
		uid = mid
	}
	// 获取当前朋友圈作者用户信息
	user, err := service.Svc.UserInfoByID(c.Request.Context(), api.GetUserID(c))
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}

	limit, offset := api.GetPage(c)
	list, err := service.Svc.MomentList(c.Request.Context(), mid, uid, limit, offset)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, map[string]any{
		"user": resource.UserBasicResource(user),
		"list": resource.MomentListResource(list.Moments, list.Users, list.Likes, list.Comments),
	})
}

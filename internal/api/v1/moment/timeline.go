package moment

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/resource"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// Timeline 我的朋友圈
// @Summary 我的朋友圈
// @Description 我的朋友圈
// @Tags 朋友圈
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param p query int false "页码"
// @success 0 {object} app.Response{data=[]model.Moment} "调用成功结构"
// @Router /moment/timeline [get]
func Timeline(c *gin.Context) {
	mid := api.GetUserID(c)
	// 获取我的用户信息
	user, err := service.Svc.UserInfoByID(c.Request.Context(), mid)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}

	limit, offset := api.GetPage(c)
	list, err := service.Svc.MomentTimeline(c.Request.Context(), mid, limit, offset)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, map[string]any{
		"user": resource.UserBasicResource(user),
		"list": resource.MomentListResource(list.Moments, list.Users, list.Likes, list.Comments),
	})
}

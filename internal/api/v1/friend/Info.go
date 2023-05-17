package friend

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"gin-chat/internal/api"
	"gin-chat/internal/resource"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// Info 获取好友信息
// @Summary 获取好友信息
// @Description 获取好友信息
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param id query int true "好友ID"
// @success 0 {object} app.Response{data=resource.FriendResponse} "调用成功结构"
// @Router /friend/info [get]
func Info(c *gin.Context) {
	fid := cast.ToInt(c.Query("id"))
	if fid == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	f, u, err := service.Svc.FriendInfo(c.Request.Context(), api.GetUserID(c), fid)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	tags, err := service.Svc.UserTagNames(c.Request.Context(), fid, f.Tags)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, resource.FriendResource(f, u, tags))
}

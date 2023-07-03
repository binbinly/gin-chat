package friend

import (
	"gin-chat/internal/api"
	"gin-chat/internal/model"
	"gin-chat/internal/resource"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
	fid := util.MustInt(c.Query("id"))
	if fid == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	uid := api.GetUserID(c)
	f, u, err := service.Svc.FriendInfo(c.Request.Context(), uid, fid)
	// 还不是好友关系
	if errors.Is(err, service.ErrFriendNotRecord) {
		f = &model.FriendModel{}
	} else if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	tags, err := service.Svc.UserTagNames(c.Request.Context(), uid, f.Tags)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, resource.FriendResource(f, u, tags))
}

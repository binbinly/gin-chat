package user

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/ecode"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// reportParams 好友举报
type reportParams struct {
	UserID   int    `json:"user_id" binding:"required,numeric" example:"1"` //用户ID
	Type     int8   `json:"type" binding:"oneof=1 2" example:"1"`           // 1=用户，2=群组
	Content  string `json:"content" binding:"required" example:"test"`      // 举报内容
	Category string `json:"category" binding:"required" example:"分类"`       // 举报分类
}

// Report 好友举报
// @Summary 好友举报
// @Description 好友举报
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body reportParams true "report"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /report [post]
func Report(c *gin.Context) {
	var req reportParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	uid := api.GetUserID(c)
	if uid == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.ReportCreate(c.Request.Context(), uid, req.UserID, req.Type, req.Category, req.Content)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

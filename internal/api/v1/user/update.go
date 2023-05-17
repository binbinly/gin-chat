package user

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// updateParams 修改用户信息
type updateParams struct {
	Avatar   string `json:"avatar" binding:"omitempty,url" example:"http://example"` // 头像
	Nickname string `json:"nickname" binding:"omitempty,max=30" example:"test"`      // 昵称
	Sign     string `json:"sign" binding:"omitempty,max=90" example:"test"`          // 签名
}

// Update 更新用户信息
// @Summary Update a user info by the user identifier
// @Description Update a user by ID
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body updateParams true "The user info"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /user/edit [post]
func Update(c *gin.Context) {
	var req updateParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}

	um := make(map[string]any)
	if req.Avatar != "" {
		um["avatar"] = req.Avatar
	}
	if req.Nickname != "" {
		um["nickname"] = req.Nickname
	}
	if req.Sign != "" {
		um["sign"] = req.Sign
	}
	if len(um) == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.UserEdit(c.Request.Context(), api.GetUserID(c), um)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

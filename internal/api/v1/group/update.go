package group

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// updateParams 修噶群组信息
type updateParams struct {
	ID     int    `json:"id" binding:"required,numeric" example:"1"`           // 群ID
	Name   string `json:"name" binding:"omitempty,max=60" example:"name"`      // 群名
	Remark string `json:"remark" binding:"omitempty,max=500" example:"remark"` // 群公告
}

// nicknameParams 修改群昵称
type nicknameParams struct {
	ID       int    `json:"id" binding:"required,numeric" example:"1"`         // 群ID
	Nickname string `json:"nickname" binding:"required,max=60" example:"name"` // 群名
}

// Update 更新群组信息
// @Summary 更新群组信息
// @Description 更新群组信息
// @Tags 群组
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param user body updateParams true "The group info"
// @success 0 {object} app.Response "调用成功结构"
// @Router /group/edit [post]
func Update(c *gin.Context) {
	var req updateParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	var err error
	if req.Remark != "" {
		err = service.Svc.GroupEditRemark(c.Request.Context(), api.GetUserID(c), req.ID, req.Remark)
	} else if req.Name != "" {
		err = service.Svc.GroupEditName(c.Request.Context(), api.GetUserID(c), req.ID, req.Name)
	} else {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

// UpdateNickname 更新群昵称
// @Summary 更新群昵称
// @Description 更新群昵称
// @Tags 群组
// @Accept  json
// @Produce  json
// @Param user body nicknameParams true "The group info"
// @success 0 {object} app.Response "调用成功结构"
// @Router /group/nickname [post]
func UpdateNickname(c *gin.Context) {
	var req nicknameParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.GroupEditUserNickname(c.Request.Context(), api.GetUserID(c), req.ID, req.Nickname)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

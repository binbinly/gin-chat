package moment

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// createParams 发布朋友圈
type createParams struct {
	Content  string `json:"content" binding:"omitempty,max=500" example:"test"`    // 内容
	Image    string `json:"image" binding:"omitempty,max=900" example:"a.jpg"`     // 图片
	Video    string `json:"video" binding:"omitempty,max=100" example:"a.mp4"`     // 视频
	Type     int8   `json:"type" binding:"required,oneof=1 2 3" example:"1"`       // 类型 1=文本 2=图文 3=视频
	Location string `json:"location" binding:"omitempty,max=100" example:"北京"`     // 地理位置
	Remind   []int  `json:"remind" example:"1,2"`                                  // 提醒用户列表
	SeeType  int8   `json:"see_type" binding:"required,oneof=1 2 3 4" example:"1"` // 可见类型 1=全部 2=私密 3=谁可见 4=谁不可见
	See      []int  `json:"see" example:"1,2"`                                     // id列表
}

// Create 发布
// @Summary 发布朋友圈
// @Description 发布朋友圈
// @Tags 朋友圈
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body createParams true "create"
// @success 0 {object} app.Response "调用成功结构"
// @Router /moment/create [post]
func Create(c *gin.Context) {
	var req createParams
	if err := api.BindJSON(c, &req); err != nil {
		app.ErrorParamInvalid(c, err)
		return
	}

	err := service.Svc.MomentCreate(c.Request.Context(), api.GetUserID(c), req.Content, req.Image, req.Video, req.Location, req.Type, req.SeeType, req.Remind, req.See)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

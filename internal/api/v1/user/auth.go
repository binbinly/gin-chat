package user

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/resource"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// loginParams 默认登录方式-用户名
type loginParams struct {
	Username string `json:"username" form:"username" binding:"required,min=4,max=18" example:"test"`   //用户名
	Password string `json:"password" form:"password" binding:"required,min=6,max=20" example:"123456"` //密码
}

// phoneLoginParams 手机号登录
type phoneLoginParams struct {
	Phone      int64  `json:"phone" form:"phone" binding:"required" example:"13333333333"`        //手机号
	VerifyCode string `json:"verify_code" form:"verify_code" binding:"required" example:"888888"` //验证码
}

// authResponse 用户登录返回
type authResponse struct {
	Token string                 `json:"token"`
	User  *resource.UserResponse `json:"user"`
}

// Login 用户名密码登录
// @Summary 用户登录接口
// @Description 通过用户名密码登录
// @Tags 用户
// @Accept json
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @success 0 {object} app.Response{data=authResponse} "调用成功结构"
// @Router /login [post]
func Login(c *gin.Context) {
	var req loginParams
	if err := api.BindJSON(c, &req); err != nil {
		app.ErrorParamInvalid(c, err)
		return
	}

	user, token, err := service.Svc.UsernameLogin(c.Request.Context(), req.Username, req.Password)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, &authResponse{
		User:  resource.UserResource(user),
		Token: token,
	})
}

// PhoneLogin 手机登录接口
// @Summary 用户登录接口
// @Description 仅限手机登录
// @Tags 用户
// @Accept json
// @Produce  json
// @Param req body phoneLoginParams true "phone"
// @success 0 {object} app.Response{data=authResponse} "调用成功结构"
// @Router /login_phone [post]
func PhoneLogin(c *gin.Context) {
	var req phoneLoginParams
	if err := api.BindJSON(c, &req); err != nil {
		app.ErrorParamInvalid(c, err)
		return
	}

	err := service.Svc.CheckVCode(c.Request.Context(), req.Phone, req.VerifyCode)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}

	user, token, err := service.Svc.UserPhoneLogin(c.Request.Context(), req.Phone)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, &authResponse{
		User:  resource.UserResource(user),
		Token: token,
	})
}

// Logout 注销登录
// @Summary 用户注销登录
// @Description 用户注销登录
// @Tags 用户
// @Accept json
// @Produce  json
// @Param Token header string true "user token"
// @success 0 {object} app.Response "调用成功结构"
// @Router /logout [get]
func Logout(c *gin.Context) {
	err := service.Svc.UserLogout(c.Request.Context(), api.GetUserID(c))
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

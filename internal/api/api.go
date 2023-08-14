package api

import (
	"errors"

	"gin-chat/internal/ecode"
	"gin-chat/internal/service"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/util"
	"github.com/gin-gonic/gin"
)

const (
	// _defaultLimit 默认分页大小
	_defaultLimit = 20
)

// BindJSON 绑定请求参数
func BindJSON(c *gin.Context, form any) error {
	if err := c.ShouldBindJSON(form); err != nil {
		logger.Debugf("[api.bind.json] param err: %v", err)
		return err
	}

	return nil
}

// GetUserID 返回用户id
func GetUserID(c *gin.Context) int {
	if c == nil {
		return 0
	}
	// uid 必须和 middleware/auth 中的 uid 命名一致
	return c.GetInt("uid")
}

// GetPage 获取分页起始偏移量
func GetPage(c *gin.Context) (int, int) {
	offset := 0
	page := util.MustInt(c.Query("p"))
	if page > 0 {
		offset = (page - 1) * _defaultLimit
	}
	return offset, _defaultLimit
}

// Error response err
func Error(err error) *errno.Error {
	switch {
	case errors.Is(err, service.ErrApplyExisted):
		return ecode.ErrApplyRepeatFailed
	case errors.Is(err, service.ErrApplyNotFound):
		return ecode.ErrApplyNotFoundFailed
	case errors.Is(err, service.ErrFriendNotFound):
		return ecode.ErrChatNotFound
	case errors.Is(err, service.ErrGroupNotFound):
		return ecode.ErrGroupNotFound
	case errors.Is(err, service.ErrGroupUserNotJoin):
		return ecode.ErrGroupNotJoin
	case errors.Is(err, service.ErrFriendNotRecord):
		return ecode.ErrFriendNotFound
	case errors.Is(err, service.ErrUserNotFound):
		return ecode.ErrUserNotFound
	case errors.Is(err, service.ErrMomentNotFound):
		return ecode.ErrMomentNotFound
	case errors.Is(err, service.ErrReportExisted):
		return ecode.ErrReportHanding
	case errors.Is(err, service.ErrGroupUserTargetNotJoin):
		return ecode.ErrGroupSelectNotJoin
	case errors.Is(err, service.ErrGroupUserExisted):
		return ecode.ErrGroupExisted
	case errors.Is(err, service.ErrUserExisted):
		return ecode.ErrUserKeyExisted
	case errors.Is(err, service.ErrUserNotMatch):
		return ecode.ErrPasswordIncorrect
	case errors.Is(err, service.ErrUserFrozen):
		return ecode.ErrUserFrozen
	case errors.Is(err, service.ErrUserTokenExpired):
		return ecode.ErrUserTokenExpired
	case errors.Is(err, service.ErrUserTokenError):
		return ecode.ErrUserTokenInvalid
	case errors.Is(err, service.ErrVerifyCodeRuleMinute):
		return ecode.ErrSendSMSMinute
	case errors.Is(err, service.ErrVerifyCodeRuleHour):
		return ecode.ErrSendSMSHour
	case errors.Is(err, service.ErrVerifyCodeRuleDay):
		return ecode.ErrSendSMSTooMany
	case errors.Is(err, service.ErrVerifyCodeNotMatch):
		return ecode.ErrVerifyCode
	case errors.Is(err, service.ErrUserOffline):
		return ecode.ErrUserOffline
	case errors.Is(err, service.ErrUserTokenEmpty):
		return ecode.ErrUserTokenEmpty
	case err != nil:
		logger.Warnf("[api] err:%v", err)
		return errno.ErrDatabase
	default:
		return nil
	}
}

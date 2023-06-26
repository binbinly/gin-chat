package ecode

import "github.com/binbinly/pkg/errno"

// nolint: golint
var (
	ErrUserNotFound          = errno.NewError(20101, "用户不存在")
	ErrPasswordIncorrect     = errno.NewError(20102, "账号或密码错误")
	ErrAreaCodeEmpty         = errno.NewError(20103, "手机区号不能为空")
	ErrPhoneEmpty            = errno.NewError(20104, "手机号不能为空")
	ErrGenVCode              = errno.NewError(20105, "生成验证码错误")
	ErrSendSMS               = errno.NewError(20106, "发送短信错误")
	ErrSendSMSMinute         = errno.NewError(20107, "一分钟限制一次哦")
	ErrSendSMSHour           = errno.NewError(20108, "触发小时级限制")
	ErrSendSMSTooMany        = errno.NewError(20109, "触发天级限制")
	ErrVerifyCode            = errno.NewError(20110, "验证码错误")
	ErrUserTokenExpired      = errno.NewError(20111, "登录过期了")
	ErrTwicePasswordNotMatch = errno.NewError(20112, "两次密码输入不一致")
	ErrUserOffline           = errno.NewError(20113, "您已离线，请重连")
	ErrUserFrozen            = errno.NewError(20114, "账号已被冻结")
	ErrUserNoSelf            = errno.NewError(20115, "不可以操作自己哦")
	ErrPhoneValid            = errno.NewError(20116, "手机号不合法")
	ErrUserKeyExisted        = errno.NewError(20117, "用户名或手机号已存在哦")
	ErrUploadImageLimit      = errno.NewError(20118, "文件太大了")
	ErrUploadNotImage        = errno.NewError(20119, "请上传文件")
	ErrUserTokenInvalid      = errno.NewError(20120, "用户凭证无效")
	ErrUserTokenEmpty        = errno.NewError(20121, "用户凭证为空")

	ErrApplyRepeatFailed   = errno.NewError(20202, "已经申请过了哦")
	ErrApplyNotFoundFailed = errno.NewError(20204, "申请未找到哦")

	ErrChatNotFound = errno.NewError(20301, "好友不存在或已被拉黑")

	ErrFriendNotFound = errno.NewError(20401, "好友没有找到哦")

	ErrGroupNotJoin       = errno.NewError(20502, "还不是群成员哦")
	ErrGroupNotFound      = errno.NewError(20503, "群聊不存在哦")
	ErrGroupExisted       = errno.NewError(20504, "已经是群成员了哦")
	ErrGroupSelectNotJoin = errno.NewError(20505, "请选择群成员哦")

	ErrMomentNotFound = errno.NewError(20601, "动态不存在哦")

	ErrReportHanding = errno.NewError(20701, "投诉正在处理中哦")
)

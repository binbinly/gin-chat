package app

import (
	"fmt"
	"time"
)

const (
	// HeaderLoginToken 登录验证 Token，Header 中传递的参数
	HeaderLoginToken = "Token"

	// HeaderSignToken 签名验证 Authorization，Header 中传递的参数
	HeaderSignToken = "Auth"

	// HeaderSignTokenDate 签名验证 Date，Header 中传递的参数
	HeaderSignTokenDate = "Auth-Date"

	// HeaderSignTokenTimeout 签名有效期为 2 分钟
	HeaderSignTokenTimeout = time.Minute * 2

	// SignSecretKey 签名秘钥
	SignSecretKey = "hWOL0E7fEH0H2cCDInLyA7WstYAwVosL"
)

const (
	// _onlinePrefix 在线key前缀
	_onlinePrefix = "user:online:"
	// _userPrefix 用户令牌标识 用于单点登录
	_userPrefix = "user:token:"
	// _historyPrefix 离线消息前缀
	_historyPrefix = "history:message:%d"
)

// BuildHistoryKey 历史消息键
func BuildHistoryKey(uid int) string {
	return fmt.Sprintf(_historyPrefix, uid)
}

// BuildOnlineKey 用户在线键
func BuildOnlineKey(uid int) string {
	return fmt.Sprintf("%s%d", _onlinePrefix, uid)
}

// BuildUserTokenKey 用户令牌键
func BuildUserTokenKey(uid int) string {
	return fmt.Sprintf("%s%d", _userPrefix, uid)
}

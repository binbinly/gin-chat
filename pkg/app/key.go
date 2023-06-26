package app

import "fmt"

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

package response

type ResCode int64

const (
	CodeSuccess = 2000 + iota
	CodeServerBusy
	CodeIllegalLogin
	CodeLoginFailure
	CodeTokenInvalid
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "请求成功",
	CodeServerBusy:   "服务繁忙",
	CodeIllegalLogin: "未登录或非法访问",
	CodeLoginFailure: "账户异地登录或令牌失效",
	CodeTokenInvalid: "无效的Token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}

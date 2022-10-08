package controller

type ResCode int64

const (
	CodeSuccess ResCode = 10000 + iota
	CodeInvalidPram
	CodeAccountExist
	CodeAccountNotExist
	CodeSeverBusy
	CodeMissAccountOrPassword
	CodeServerNotFound
	CodeServerPing

	CodeTokenException
	CodeInvalidToken
	CodeNotLogin
	CodeRefreshTokenFail
)

var resMsg = map[ResCode]string{
	CodeSuccess:               "success",
	CodeInvalidPram:           "参数无效",
	CodeAccountExist:          "该账户已存在",
	CodeAccountNotExist:       "该账户不存在",
	CodeSeverBusy:             "服务繁忙",
	CodeMissAccountOrPassword: "用户名或密码错误",
	CodeServerNotFound:        "404 Not Found",
	CodeServerPing:            "pong",
	CodeTokenException:        "token 异常",

	CodeInvalidToken:     "无效Token",
	CodeNotLogin:         "Token 未登录",
	CodeRefreshTokenFail: "Token 刷新失败",
}

func (res ResCode) Msg() string {
	msg, ok := resMsg[res]
	if !ok {
		return resMsg[CodeSeverBusy]
	}
	return msg
}

package errno

import "net/http"

var (
	// OK
	OK = NewError(0, "OK")

	// 服务级错误码
	ErrServer      = NewError(10101, http.StatusText(http.StatusInternalServerError))
	Err404         = NewError(10102, http.StatusText(http.StatusNotFound))
	ErrManyRequest = NewError(10103, "Too many requests")

	ErrParam     = NewError(10002, "参数有误")
	ErrSignParam = NewError(10003, "签名参数有误")

	// 模块级错误码 - 用户模块
	ErrUser        = NewError(20101, "非法用户")
	ErrUserCaptcha = NewError(20102, "用户验证码有误")

	// ...
)

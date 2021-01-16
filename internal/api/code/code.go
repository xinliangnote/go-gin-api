package code

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

var (
	// OK
	OK = errno.NewError(1, "OK")

	// 服务级错误码
	ErrServer      = errno.NewError(10101, http.StatusText(http.StatusInternalServerError))
	ErrManyRequest = errno.NewError(10102, "Too many requests")

	ErrParam     = errno.NewError(10110, "参数有误")
	ErrSignParam = errno.NewError(10111, "缺少签名")
	ErrSign      = errno.NewError(10112, "签名有误")

	// 模块级错误码 - 用户模块
	ErrUser       = errno.NewError(20101, "非法用户")
	ErrUserCreate = errno.NewError(20102, "创建用户失败")
	ErrUserUpdate = errno.NewError(20103, "更新用户失败")
	ErrUserSearch = errno.NewError(20104, "查询用户失败")
	ErrUserHTTP   = errno.NewError(20105, "调用他方接口失败")

	// ...
)

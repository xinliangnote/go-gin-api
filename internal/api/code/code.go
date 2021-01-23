package code

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

var (
	// OK
	OK = errno.NewError(http.StatusOK, 1, "OK")

	// 服务级错误码
	ErrServer        = errno.NewError(http.StatusInternalServerError, 10101, http.StatusText(http.StatusInternalServerError))
	ErrManyRequest   = errno.NewError(http.StatusTooManyRequests, 10102, http.StatusText(http.StatusTooManyRequests))
	ErrParamBind     = errno.NewError(http.StatusBadRequest, 10103, "参数信息有误")
	ErrAuthorization = errno.NewError(http.StatusUnauthorized, 10104, "签名信息有误")

	// 模块级错误码 - 用户模块
	ErrUser       = errno.NewError(http.StatusBadRequest, 20101, "非法用户")
	ErrUserName   = errno.NewError(http.StatusBadRequest, 20102, "账号不能为空")
	ErrUserCreate = errno.NewError(http.StatusBadRequest, 20103, "创建用户失败")
	ErrUserUpdate = errno.NewError(http.StatusBadRequest, 20104, "更新用户失败")
	ErrUserSearch = errno.NewError(http.StatusBadRequest, 20105, "查询用户失败")
	ErrUserHTTP   = errno.NewError(http.StatusBadRequest, 20106, "调用他方接口失败")

	// ...
)

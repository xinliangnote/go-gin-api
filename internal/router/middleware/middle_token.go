package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
)

func (m *middleware) Token(ctx core.Context) (userId int64, userName string, err errno.Error) {
	token := ctx.GetHeader(configs.HeaderLoginToken)
	if token == "" {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("Header 中缺少 Token 参数"))

		return
	}

	if !m.cache.Exists(configs.RedisKeyPrefixLoginUser + token) {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(errors.New("请先登录"))

		return
	}

	cacheData, cacheErr := m.cache.Get(configs.RedisKeyPrefixLoginUser+token, cache.WithTrace(ctx.Trace()))
	if cacheErr != nil {
		err = errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(cacheErr)

		return
	}

	type userInfo struct {
		Id       int64  `json:"id"`       // 用户ID
		Username string `json:"username"` // 用户名
	}

	var userData userInfo

	jsonErr := json.Unmarshal([]byte(cacheData), &userData)
	if jsonErr != nil {
		errno.NewError(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithErr(jsonErr)

		return
	}

	userId = userData.Id
	userName = userData.Username

	return
}

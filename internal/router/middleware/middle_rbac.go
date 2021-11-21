package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/services/admin"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
	"github.com/xinliangnote/go-gin-api/pkg/urltable"
)

func (m *middleware) RBAC() core.HandlerFunc {
	return func(c core.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			c.AbortWithError(errno.NewError(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(errors.New("Header 中缺少 Token 参数")),
			)
			return
		}

		if !m.cache.Exists(configs.RedisKeyPrefixLoginUser + token) {
			c.AbortWithError(errno.NewError(
				http.StatusUnauthorized,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithErr(errors.New("请先登录 1")),
			)
			return
		}

		if !m.cache.Exists(configs.RedisKeyPrefixLoginUser + token + ":action") {
			c.AbortWithError(errno.NewError(
				http.StatusUnauthorized,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithErr(errors.New("请先登录 2")),
			)
			return
		}

		actionData, err := m.cache.Get(configs.RedisKeyPrefixLoginUser+token+":action", redis.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusUnauthorized,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithErr(err),
			)
			return
		}

		var actions []admin.MyActionData
		err = json.Unmarshal([]byte(actionData), &actions)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(err),
			)
			return
		}

		if len(actions) > 0 {
			table := urltable.NewTable()
			for _, v := range actions {
				_ = table.Append(v.Method + v.Api)
			}

			if pattern, _ := table.Mapping(c.Method() + c.Path()); pattern == "" {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.RBACError,
					code.Text(code.RBACError)).WithErr(errors.New(c.Method() + c.Path() + " 未进行 RBAC 授权")),
				)
				return
			}
		}

	}
}

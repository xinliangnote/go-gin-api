package middleware

import "github.com/xinliangnote/go-gin-api/internal/pkg/core"

func (m *middleware) DisableLog() core.HandlerFunc {
	return func(c core.Context) {
		core.DisableTrace(c)
	}
}

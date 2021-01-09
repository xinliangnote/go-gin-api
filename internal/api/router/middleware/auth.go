package middleware

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/token"
)

func AuthHandler(ctx core.Context) (userId int, userName string, err errno.Error) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		err = code.ErrSignParam
		return
	}

	cfg := configs.Get().JWT
	claims, errParse := token.New(cfg.Secret).Parse(auth)
	if errParse != nil {
		err = code.ErrSignParam
		return
	}

	userId = claims.UserID
	if userId <= 0 {
		err = code.ErrSignParam
		return
	}
	userName = claims.UserName
	return
}

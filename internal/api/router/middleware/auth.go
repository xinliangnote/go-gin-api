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
		err = code.ErrAuthorization
		return
	}

	cfg := configs.Get().JWT
	claims, errParse := token.New(cfg.Secret).Parse(auth)
	if errParse != nil {
		err = code.ErrAuthorization
		return
	}

	userId = claims.UserID
	if userId <= 0 {
		err = code.ErrAuthorization
		return
	}
	userName = claims.UserName
	return
}

package middleware

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/errno"
	"github.com/xinliangnote/go-gin-api/internal/pkg/token"
)

func AuthHandler(ctx core.Context) (userId int, userName string, err errno.Error) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		err = errno.ErrSignParam
		return
	}

	claims, errParse := token.Parse(auth)
	if errParse != nil {
		err = errno.ErrSignParam
		return
	}

	userId = claims.UserID
	if userId <= 0 {
		err = errno.ErrSignParam
		return
	}
	userName = claims.UserName
	return
}

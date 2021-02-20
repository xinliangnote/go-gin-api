package user_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/user_demo_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (u *userSer) GetUserByUserName(ctx core.Context, username string) (user *user_demo_repo.UserDemo, err error) {
	user, err = u.userRepo.GetUserByUserName(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

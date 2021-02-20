package user_service

import "github.com/xinliangnote/go-gin-api/internal/pkg/core"

func (u *userSer) UpdateNickNameByID(ctx core.Context, id uint, username string) (err error) {
	err = u.userRepo.UpdateNickNameByID(ctx, id, username)
	if err != nil {
		return nil
	}
	return nil
}

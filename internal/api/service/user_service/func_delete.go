package user_service

import "github.com/xinliangnote/go-gin-api/internal/pkg/core"

func (u *userSer) Delete(ctx core.Context, id uint) (err error) {
	err = u.userRepo.Delete(ctx, id)
	if err != nil {
		return nil
	}
	return nil
}

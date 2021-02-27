package user_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/user_demo_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ UserService = (*userSer)(nil)

type UserService interface {
	// i 为了避免被其他包实现
	i()

	Create(ctx core.Context, user *CreateUserInfo) (id int32, err error)
	UpdateNickNameByID(ctx core.Context, id int32, username string) (err error)
	GetUserByUserName(ctx core.Context, username string) (user *user_demo_repo.UserDemo, err error)
	Delete(ctx core.Context, id int32) (err error)
}

type userSer struct {
	db    db.Repo
	cache cache.Repo
}

func NewUserService(db db.Repo, cache cache.Repo) UserService {
	return &userSer{
		db:    db,
		cache: cache,
	}
}

func (u *userSer) i() {}

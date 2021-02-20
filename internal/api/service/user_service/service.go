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

	Create(ctx core.Context, user *CreateUserInfo) (id uint, err error)
	UpdateNickNameByID(ctx core.Context, id uint, username string) (err error)
	GetUserByUserName(ctx core.Context, username string) (user *user_demo_repo.UserDemo, err error)
	Delete(ctx core.Context, id uint) (err error)
}

type userSer struct {
	db       db.Repo
	cache    cache.Repo
	userRepo user_demo_repo.UserRepo
}

func NewUserService(db db.Repo, cache cache.Repo) UserService {
	userRepo := user_demo_repo.NewUserRepo(db)
	return &userSer{
		db:       db,
		cache:    cache,
		userRepo: userRepo,
	}
}

func (u *userSer) i() {}

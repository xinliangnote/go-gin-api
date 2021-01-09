package user_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/model/user_model"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/cache_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/user_demo_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

var _ UserService = (*userSer)(nil)

type UserService interface {
	// i 为了避免被其他包实现
	i()

	Create(ctx core.Context, user *user_model.CreateRequest) (id uint, err error)
	UpdateNickNameByID(ctx core.Context, id uint, username string) (err error)
	GetUserByUserName(ctx core.Context, username string) (user *user_model.UserDemo, err error)
}

type userSer struct {
	db       db_repo.Repo
	cache    cache_repo.Repo
	userRepo user_demo_repo.UserRepo
}

func NewUserService(db db_repo.Repo, cache cache_repo.Repo) UserService {
	userRepo := user_demo_repo.NewUserRepo(db)
	return &userSer{
		db:       db,
		cache:    cache,
		userRepo: userRepo,
	}
}

func (u *userSer) i() {}

func (u *userSer) Create(ctx core.Context, user *user_model.CreateRequest) (id uint, err error) {
	create := user_model.UserDemo{
		UserName: user.UserName,
		NickName: user.NickName,
		Mobile:   user.Mobile,
	}

	id, err = u.userRepo.Create(ctx, create)
	if err != nil {
		return 0, err
	}
	return
}

func (u *userSer) UpdateNickNameByID(ctx core.Context, id uint, username string) (err error) {
	err = u.userRepo.UpdateNickNameByID(ctx, id, username)
	if err != nil {
		return nil
	}
	return nil
}

func (u *userSer) GetUserByUserName(ctx core.Context, username string) (user *user_model.UserDemo, err error) {
	user, err = u.userRepo.GetUserByUserName(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

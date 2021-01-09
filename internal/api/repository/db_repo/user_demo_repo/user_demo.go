package user_demo_repo

import (
	"github.com/xinliangnote/go-gin-api/internal/api/model/user_model"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var _ UserRepo = (*userRepo)(nil)

type UserRepo interface {
	// i 为了避免被其他包实现
	i()
	Create(ctx core.Context, user user_model.UserDemo) (id uint, err error)
	UpdateNickNameByID(ctx core.Context, id uint, username string) (err error)
	GetUserByUserName(ctx core.Context, username string) (*user_model.UserDemo, error)
	getUserByID(ctx core.Context, id uint) (*user_model.UserDemo, error)
}

type userRepo struct {
	db db_repo.Repo
}

func NewUserRepo(db db_repo.Repo) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) i() {}

func (u *userRepo) Create(ctx core.Context, user user_model.UserDemo) (id uint, err error) {
	err = u.db.GetDbW().WithContext(ctx).Create(&user).Error
	if err != nil {
		return 0, errors.Wrap(err, "[user_repo] create user err")
	}
	return user.Id, nil
}

func (u *userRepo) getUserByID(ctx core.Context, id uint) (*user_model.UserDemo, error) {
	data := new(user_model.UserDemo)
	err := u.db.GetDbR().WithContext(ctx).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[user_demo] get user data err")
	}
	return data, nil
}

func (u *userRepo) UpdateNickNameByID(ctx core.Context, id uint, nickname string) (err error) {
	user, err := u.getUserByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[user_demo] update user data err")
	}
	return u.db.GetDbW().WithContext(ctx).Model(user).Update("nick_name", nickname).Error
}

func (u *userRepo) GetUserByUserName(ctx core.Context, username string) (*user_model.UserDemo, error) {
	data := new(user_model.UserDemo)
	err := u.db.GetDbR().
		WithContext(ctx).
		Select([]string{"id", "user_name", "nick_name", "mobile"}).
		Where("user_name = ?", username).
		First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[user_demo] get user data err")
	}
	return data, nil
}

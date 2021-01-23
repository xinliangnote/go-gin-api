package user_demo_repo

import (
	"github.com/xinliangnote/go-gin-api/internal/api/model/user_model"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"

	"github.com/pkg/errors"
)

var _ UserRepo = (*userRepo)(nil)

type UserRepo interface {
	// i 为了避免被其他包实现
	i()
	Create(ctx core.Context, user user_model.UserDemo) (id uint, err error)
	UpdateNickNameByID(ctx core.Context, id uint, username string) (err error)
	GetUserByUserName(ctx core.Context, username string) (*user_model.UserDemo, error)
	Delete(ctx core.Context, id uint) (err error)
	getUserByID(ctx core.Context, id uint) (*user_model.UserDemo, error)
}

type userRepo struct {
	db db.Repo
}

func NewUserRepo(db db.Repo) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) i() {}

func (u *userRepo) Create(ctx core.Context, user user_model.UserDemo) (id uint, err error) {
	err = u.db.GetDbW().WithContext(ctx.RequestContext()).Create(&user).Error
	if err != nil {
		return 0, errors.Wrap(err, "[user_repo] create user err")
	}
	return user.Id, nil
}

func (u *userRepo) getUserByID(ctx core.Context, id uint) (*user_model.UserDemo, error) {
	data := new(user_model.UserDemo)
	err := u.db.GetDbR().WithContext(ctx.RequestContext()).First(data, id).Where("is_deleted = ?", -1).Error
	if err != nil {
		return nil, errors.Wrap(err, "[user_demo] get user data err")
	}
	return data, nil
}

func (u *userRepo) UpdateNickNameByID(ctx core.Context, id uint, nickname string) (err error) {
	user, err := u.getUserByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[user_demo] update user data err")
	}
	return u.db.GetDbW().WithContext(ctx.RequestContext()).Model(user).Update("nick_name", nickname).Error
}

func (u *userRepo) Delete(ctx core.Context, id uint) (err error) {
	user, err := u.getUserByID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[user_demo] update user data err")
	}
	return u.db.GetDbW().WithContext(ctx.RequestContext()).Model(user).Update("is_deleted", 1).Error
}

func (u *userRepo) GetUserByUserName(ctx core.Context, username string) (*user_model.UserDemo, error) {
	data := new(user_model.UserDemo)
	err := u.db.GetDbR().
		WithContext(ctx.RequestContext()).
		Select([]string{"id", "user_name", "nick_name", "mobile"}).
		Where("user_name = ? and is_deleted = ?", username, -1).
		First(data).Error
	if err != nil {
		return nil, errors.Wrap(err, "[user_demo] get user data err")
	}
	return data, nil
}

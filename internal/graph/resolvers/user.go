package resolvers

import (
	"context"

	"github.com/xinliangnote/go-gin-api/internal/graph/model"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
)

func (r *queryResolver) BySex(ctx context.Context, sex string) ([]*model.User, error) {
	if sex == "" {
		return nil, errors.New("sex required")
	}

	//模拟数据
	var users []*model.User
	users = append(users, &model.User{ID: "1", Name: "Tom", Sex: sex, Mobile: "13266666666"})
	users = append(users, &model.User{ID: "1", Name: "Jack", Sex: sex, Mobile: "13288888888"})

	return users, nil
}

func (r *mutationResolver) UpdateUserMobile(ctx context.Context, data model.UpdateUserMobileInput) (*model.User, error) {
	if data.ID == "" {
		return nil, errors.New("id required")
	}

	if data.Mobile == "" {
		return nil, errors.New("mobile required")
	}

	//模拟数据
	user := new(model.User)
	user.ID = data.ID
	user.Mobile = data.Mobile
	user.Sex = "男"
	user.Name = "Jack"

	//操作数据库
	//userData, err := r.userService.GetUserByUserName(r.getCoreContextByCtx(ctx), "test_user")
	//if err != nil {
	//	return nil, err
	//}

	return user, nil
}

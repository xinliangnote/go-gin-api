package menu_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_action_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/menu_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, menuData *CreateMenuData) (id int32, err error)
	Modify(ctx core.Context, id int32, menuData *UpdateMenuData) (err error)
	List(ctx core.Context, searchData *SearchData) (listData []*menu_repo.Menu, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	UpdateSort(ctx core.Context, id int32, sort int32) (err error)
	Delete(ctx core.Context, id int32) (err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *menu_repo.Menu, err error)

	CreateAction(ctx core.Context, menuActionData *CreateMenuActionData) (id int32, err error)
	ListAction(ctx core.Context, searchListActionData *SearchListActionData) (listData []*menu_action_repo.MenuAction, err error)
	DeleteAction(ctx core.Context, id int32) (err error)
}

type service struct {
	db    db.Repo
	cache cache.Repo
}

func New(db db.Repo, cache cache.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}

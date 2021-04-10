package admin_service

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ Service = (*service)(nil)

// 定义缓存前缀
var cacheKeyPrefix = configs.ProjectName() + ":admin:"

type Service interface {
	i()
	CacheKeyPrefix() (pre string)

	Create(ctx core.Context, authorizedData *CreateAdminData) (id int32, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*admin_repo.Admin, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	Delete(ctx core.Context, id int32) (err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *admin_repo.Admin, err error)
	ResetPassword(ctx core.Context, id int32) (err error)
	ModifyPassword(ctx core.Context, id int32, newPassword string) (err error)
	ModifyPersonalInfo(ctx core.Context, id int32, modifyData *ModifyData) (err error)
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

func (s *service) CacheKeyPrefix() (pre string) {
	pre = cacheKeyPrefix
	return
}

package authorized_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_api_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, authorizedData *CreateAuthorizedData) (id int32, err error)
	List(ctx core.Context, searchData *SearchData) (listData []*authorized_repo.Authorized, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*authorized_repo.Authorized, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	Delete(ctx core.Context, id int32) (err error)
	Detail(ctx core.Context, id int32) (info *authorized_repo.Authorized, err error)
	DetailByKey(ctx core.Context, key string) (data *CacheAuthorizedData, err error)

	CreateAPI(ctx core.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int32, err error)
	ListAPI(ctx core.Context, searchAPIData *SearchAPIData) (listData []*authorized_api_repo.AuthorizedApi, err error)
	DeleteAPI(ctx core.Context, id int32) (err error)
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

package cron

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/cron"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/cron_task_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, createData *CreateCronTaskData) (id int32, err error)
	Modify(ctx core.Context, id int32, modifyData *ModifyCronTaskData) (err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*cron_task_repo.CronTask, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	Execute(ctx core.Context, id int32) (err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *cron_task_repo.CronTask, err error)
}

type service struct {
	db         db.Repo
	cache      redis.Repo
	cronServer cron.Server
}

func New(db db.Repo, cache redis.Repo, cron cron.Server) Service {
	return &service{
		db:         db,
		cache:      cache,
		cronServer: cron,
	}
}

func (s *service) i() {}

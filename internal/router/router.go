package router

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/pkg/grpc"
	"github.com/xinliangnote/go-gin-api/internal/pkg/metrics"
	"github.com/xinliangnote/go-gin-api/internal/pkg/notify"
	"github.com/xinliangnote/go-gin-api/internal/router/middleware"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type resource struct {
	mux     core.Mux
	logger  *zap.Logger
	db      db.Repo
	cache   cache.Repo
	grpConn grpc.ClientConn
	middles middleware.Middleware
}

func NewHTTPMux(logger *zap.Logger, db db.Repo, cache cache.Repo, grpConn grpc.ClientConn) (core.Mux, error) {

	if logger == nil {
		return nil, errors.New("logger required")
	}

	mux, err := core.New(logger,
		core.WithEnableOpenBrowser("http://127.0.0.1:9999"),
		core.WithEnableCors(),
		core.WithEnableRate(),
		core.WithPanicNotify(notify.OnPanicNotify),
		core.WithRecordMetrics(metrics.RecordMetrics),
	)

	if err != nil {
		panic(err)
	}

	r := new(resource)
	r.mux = mux
	r.logger = logger
	r.db = db
	r.cache = cache
	r.grpConn = grpConn
	r.middles = middleware.New(logger, cache, db)

	// 设置 WEB 路由
	setWebRouter(r)

	// 设置 API 路由
	setApiRouter(r)

	// 设置 GraphQL 路由
	setGraphQLRouter(r)

	return mux, nil
}

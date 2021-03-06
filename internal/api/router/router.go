package router

import (
	"github.com/xinliangnote/go-gin-api/internal/api/controller/demo_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/user_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/router/middleware"
	"github.com/xinliangnote/go-gin-api/internal/graph/handler"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/pkg/grpc"
	"github.com/xinliangnote/go-gin-api/internal/pkg/metrics"
	"github.com/xinliangnote/go-gin-api/internal/pkg/notify"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewHTTPMux(logger *zap.Logger, db db.Repo, cache cache.Repo, grpConn grpc.ClientConn) (core.Mux, error) {

	if logger == nil {
		return nil, errors.New("logger required")
	}

	mux, err := core.New(logger,
		core.WithEnableCors(),
		core.WithEnableRate(),
		core.WithPanicNotify(notify.OnPanicNotify),
		core.WithRecordMetrics(metrics.RecordMetrics),
	)

	if err != nil {
		panic(err)
	}

	// 中间件
	middles := middleware.New(logger, cache)

	// graphQL 控制器
	gqlHandler := handler.New(logger, db, cache)

	gql := mux.Group("/graphql")
	{
		gql.GET("", gqlHandler.Playground())
		gql.POST("/query", gqlHandler.Query())
	}

	// demo 控制器
	demoHandler := demo_handler.New(logger, db, cache, grpConn)
	demo := mux.Group("/demo", core.WrapAuthHandler(middles.Jwt)) // 使用 jwt 验证
	{
		// 为了演示 Trace ，增加了一些看起来无意义的调试信息和 SQL 信息。
		demo.GET("/trace", demoHandler.Trace())

		// 模拟数据
		demo.GET("get/:name", core.AliasForRecordMetrics("/demo/get"), demoHandler.Get())
		demo.POST("post", demoHandler.Post())
	}

	demoNoAuth := mux.Group("/auth") // 不使用 jwt 验证
	{
		demoNoAuth.POST("/get", demoHandler.Auth())
	}

	// user 控制器
	userHandler := user_handler.New(logger, db, cache)
	user := mux.Group("/user", core.WrapAuthHandler(middles.Jwt))
	{
		user.POST("/create", userHandler.Create())
		user.PUT("/update", userHandler.UpdateNickNameByID())
		user.PATCH("/delete/:id", userHandler.Delete())
		user.GET("/info/:username", core.AliasForRecordMetrics("/user/info"), userHandler.Detail())
	}

	return mux, nil
}

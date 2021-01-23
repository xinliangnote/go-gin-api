package router

import (
	"github.com/xinliangnote/go-gin-api/internal/api/controller/demo"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/user_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/router/middleware/auth"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/pkg/metrics"
	"github.com/xinliangnote/go-gin-api/internal/pkg/notify"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewHTTPMux(logger *zap.Logger, db db.Repo, cache cache.Repo) (core.Mux, error) {

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

	demoHandler := demo.NewDemo(logger, db, cache)
	userHandler := user_handler.NewUserDemo(logger, db, cache)

	// user_demo CURD
	user := mux.Group("/user", core.WrapAuthHandler(auth.AuthHandler))
	{
		user.POST("/create", userHandler.Create())
		user.PUT("/update", userHandler.UpdateNickNameByID())
		user.PATCH("/delete/:id", userHandler.Delete())
		user.GET("/info/:username", core.AliasForRecordMetrics("/user/info"), userHandler.Detail())
	}

	// auth
	a := mux.Group("/auth")
	{
		a.POST("/get", demoHandler.Auth())
	}

	// demo
	d := mux.Group("/demo", core.WrapAuthHandler(auth.AuthHandler)) //使用 auth 验证
	{
		// 为了演示 Trace ，增加了一些看起来无意义的调试信息和 SQL 信息。
		d.GET("/trace", demoHandler.Trace())

		// 模拟数据
		d.GET("get/:name", core.AliasForRecordMetrics("/demo/get"), demoHandler.Get())
		d.POST("post", demoHandler.Post())
	}

	return mux, nil
}

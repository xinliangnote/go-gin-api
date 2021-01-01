package router

import (
	"github.com/xinliangnote/go-gin-api/internal/api/controller/demo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/metrics"
	"github.com/xinliangnote/go-gin-api/internal/pkg/notify"
	"github.com/xinliangnote/go-gin-api/internal/router/middleware"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewHTTPMux(logger *zap.Logger) (core.Mux, error) {

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

	demoHandler := demo.NewDemo(logger)

	u := mux.Group("/user")
	{
		u.POST("/login", demoHandler.Login())
	}

	d := mux.Group("/demo", core.WrapAuthHandler(middleware.AuthHandler)) //使用 auth 验证
	{
		d.GET("user/:name", demoHandler.User())

		// 模拟数据
		d.GET("get/:name", demoHandler.Get(), core.DisableJournal)
		d.POST("post", demoHandler.Post(), core.DisableJournal)

		// 测试加密性能
		d.GET("/rsa/test", demoHandler.RsaTest())
		d.GET("/aes/test", demoHandler.AesTest())
		d.GET("/md5/test", demoHandler.MD5Test())
	}

	return mux, nil
}

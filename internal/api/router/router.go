package router

import (
	"github.com/xinliangnote/go-gin-api/internal/api/controller/demo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/notify"

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
		core.WithPanicNotify(notify.Email),
	)

	if err != nil {
		panic(err)
	}

	//设置路由中间件
	//engine.Use(exception.SetUp(), jaeger.SetUp())

	demoHandler := demo.NewDemo(logger)

	d := mux.Group("/demo")
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

	// 测试链路追踪
	//mux.GET("/jaeger_test", jaeger_conn.JaegerTest)

	return mux, nil
}

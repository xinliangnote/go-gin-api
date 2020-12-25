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
		core.WithDisablePProf(),
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
		d.GET("user", demoHandler.User())

		// 模拟数据
		d.GET("get/:name", demoHandler.Get(), core.DisableJournal)
		d.POST("post", demoHandler.Post(), core.DisableJournal)

	}

	// 测试链路追踪
	//mux.GET("/jaeger_test", jaeger_conn.JaegerTest)

	// 测试加密性能
	//TestRouter := mux.Group("/test")
	//{
	//	// 测试 MD5 组合 的性能
	//	TestRouter.GET("/md5", core.Handle(test.Md5Test))
	//
	//	// 测试 AES 的性能
	//	TestRouter.GET("/aes", core.Handle(test.AesTest))
	//
	//	// 测试 RSA 的性能
	//	TestRouter.GET("/rsa", core.Handle(test.RsaTest))
	//}

	return mux, nil
}

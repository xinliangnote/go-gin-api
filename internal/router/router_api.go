package router

import (
	"github.com/xinliangnote/go-gin-api/internal/api/controller/demo_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/user_handler"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func setApiRouter(r *resource) {
	// demo 控制器
	demoHandler := demo_handler.New(r.logger, r.db, r.cache, r.grpConn)
	demo := r.mux.Group("/demo", core.WrapAuthHandler(r.middles.Jwt)) // 使用 jwt 验证
	{
		// 为了演示 Trace ，增加了一些看起来无意义的调试信息和 SQL 信息。
		demo.GET("/trace", demoHandler.Trace())

		// 模拟数据
		demo.GET("get/:name", core.AliasForRecordMetrics("/demo/get"), demoHandler.Get())
		demo.POST("post", demoHandler.Post())
	}

	demoNoAuth := r.mux.Group("/auth") // 不使用 jwt 验证
	{
		demoNoAuth.POST("/get", demoHandler.Auth())
	}

	// user 控制器
	userHandler := user_handler.New(r.logger, r.db, r.cache)
	user := r.mux.Group("/user", core.WrapAuthHandler(r.middles.Jwt))
	{
		user.POST("/create", userHandler.Create())
		user.PUT("/update", userHandler.UpdateNickNameByID())
		user.PATCH("/delete/:id", userHandler.Delete())
		user.GET("/info/:username", core.AliasForRecordMetrics("/user/info"), userHandler.Detail())
	}
}

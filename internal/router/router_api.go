package router

import (
	"github.com/xinliangnote/go-gin-api/internal/api/controller/admin_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/authorized_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/demo_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/tool_handler"
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

	// authorized
	authorizedHandler := authorized_handler.New(r.logger, r.db, r.cache)

	// admin
	adminHandler := admin_handler.New(r.logger, r.db, r.cache)

	// 登录
	login := r.mux.Group("/login", r.middles.Signature())
	{
		login.POST("/web", adminHandler.Login())
	}

	// api
	api := r.mux.Group("/api", core.WrapAuthHandler(r.middles.Token), r.middles.Signature())
	{
		api.POST("/authorized", authorizedHandler.Create())
		api.GET("/authorized", authorizedHandler.List())
		api.PATCH("/authorized/used", authorizedHandler.UpdateUsed())
		api.DELETE("/authorized/:id", authorizedHandler.Delete())

		api.POST("/authorized_api", authorizedHandler.CreateAPI())
		api.GET("/authorized_api", authorizedHandler.ListAPI())
		api.DELETE("/authorized_api/:id", authorizedHandler.DeleteAPI())

		api.POST("/admin", adminHandler.Create())
		api.GET("/admin", adminHandler.List())
		api.PATCH("/admin/used", adminHandler.UpdateUsed())
		api.PATCH("/admin/reset_password/:id", adminHandler.ResetPassword())
		api.DELETE("/admin/:id", adminHandler.Delete())
		api.POST("/admin/logout", adminHandler.Logout())
		api.PATCH("/admin/modify_password", adminHandler.ModifyPassword())
		api.GET("/admin/info", adminHandler.Detail())
		api.PATCH("/admin/modify_personal_info", adminHandler.ModifyPersonalInfo())

		// tool
		toolHandler := tool_handler.New(r.logger, r.db, r.cache)
		api.GET("/tool/hashids/encode/:id", toolHandler.HashIdsEncode())
		api.GET("/tool/hashids/decode/:id", toolHandler.HashIdsDecode())
	}
}

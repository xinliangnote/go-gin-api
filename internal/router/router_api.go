package router

import (
	"github.com/xinliangnote/go-gin-api/internal/api/controller/admin_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/authorized_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/config_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/menu_handler"
	"github.com/xinliangnote/go-gin-api/internal/api/controller/tool_handler"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func setApiRouter(r *resource) {

	// admin
	adminHandler := admin_handler.New(r.logger, r.db, r.cache)

	// login
	login := r.mux.Group("/login", r.middles.Signature())
	{
		login.POST("/web", adminHandler.Login())
	}

	// api
	api := r.mux.Group("/api", core.WrapAuthHandler(r.middles.Token), r.middles.Signature())
	{
		// authorized
		authorizedHandler := authorized_handler.New(r.logger, r.db, r.cache)
		api.POST("/authorized", authorizedHandler.Create())
		api.GET("/authorized", authorizedHandler.List())
		api.PATCH("/authorized/used", authorizedHandler.UpdateUsed())
		api.DELETE("/authorized/:id", core.AliasForRecordMetrics("/api/authorized/info"), authorizedHandler.Delete())

		api.POST("/authorized_api", authorizedHandler.CreateAPI())
		api.GET("/authorized_api", authorizedHandler.ListAPI())
		api.DELETE("/authorized_api/:id", core.AliasForRecordMetrics("/api/authorized_api/info"), authorizedHandler.DeleteAPI())

		api.POST("/admin", adminHandler.Create())
		api.GET("/admin", adminHandler.List())
		api.PATCH("/admin/used", adminHandler.UpdateUsed())
		api.PATCH("/admin/reset_password/:id", adminHandler.ResetPassword())
		api.DELETE("/admin/:id", adminHandler.Delete())
		api.POST("/admin/logout", adminHandler.Logout())
		api.PATCH("/admin/modify_password", adminHandler.ModifyPassword())
		api.GET("/admin/info", adminHandler.Detail())
		api.PATCH("/admin/modify_personal_info", adminHandler.ModifyPersonalInfo())

		// menu
		menuHandler := menu_handler.New(r.logger, r.db, r.cache)
		api.POST("/menu", menuHandler.Create())
		api.GET("/menu", menuHandler.List())
		api.GET("/menu/:id", menuHandler.Detail())
		api.PATCH("/menu/used", menuHandler.UpdateUsed())
		api.DELETE("/menu/:id", menuHandler.Delete())

		// tool
		toolHandler := tool_handler.New(r.logger, r.db, r.cache)
		api.GET("/tool/hashids/encode/:id", toolHandler.HashIdsEncode())
		api.GET("/tool/hashids/decode/:id", toolHandler.HashIdsDecode())
		api.POST("/tool/cache/search", toolHandler.SearchCache())
		api.PATCH("/tool/cache/clear", toolHandler.ClearCache())
		api.GET("/tool/data/dbs", toolHandler.Dbs())
		api.POST("/tool/data/tables", toolHandler.Tables())
		api.POST("/tool/data/mysql", toolHandler.SearchMySQL())

		// config
		configHandler := config_handler.New(r.logger, r.db, r.cache)
		api.PATCH("/config/email", configHandler.Email())

	}
}

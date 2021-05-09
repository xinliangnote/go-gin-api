package router

import (
	"github.com/xinliangnote/go-gin-api/internal/web/controller/admin_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/authorized_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/config_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/dashboard_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/generator_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/index_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/install_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/tool_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/upgrade_handler"
)

func setWebRouter(r *resource) {

	installHandler := install_handler.New(r.logger)
	indexHandler := index_handler.New(r.logger, r.db, r.cache)
	dashboardHandler := dashboard_handler.New(r.logger, r.db, r.cache)
	generatorHandler := generator_handler.New(r.logger, r.db, r.cache)
	configInfoHandler := config_handler.New(r.logger, r.db, r.cache)
	authorizedHandler := authorized_handler.New(r.logger, r.db, r.cache)
	toolHandler := tool_handler.New(r.logger, r.db, r.cache)
	adminHandler := admin_handler.New(r.logger, r.db, r.cache)
	upgradeHandler := upgrade_handler.New(r.logger, r.db, r.cache)

	// 无需记录日志，无需 RBAC 权限验证
	notRBAC := r.mux.Group("", r.middles.DisableLog())
	{
		// 首页
		notRBAC.GET("", indexHandler.View())

		// 仪表盘
		notRBAC.GET("/dashboard", dashboardHandler.View())

		// 安装
		notRBAC.GET("/install", installHandler.View())
		notRBAC.POST("/install/execute", installHandler.Execute())
		notRBAC.POST("/install/restart", installHandler.Restart())

		// 管理员
		notRBAC.GET("/login", adminHandler.LoginView())
		notRBAC.GET("/admin/modify_password", adminHandler.ModifyPasswordView())
		notRBAC.GET("/admin/modify_info", adminHandler.ModifyInfoView())
	}

	// 无需记录日志，需要 RBAC 权限验证
	web := r.mux.Group("", r.middles.DisableLog())
	{
		// 配置信息
		web.GET("/config/email", configInfoHandler.EmailView())
		web.GET("/config/code", configInfoHandler.CodeView())

		// 代码生成器
		web.GET("/generator/gorm", generatorHandler.GormView())
		web.POST("/generator/gorm/execute", generatorHandler.GormExecute())

		web.GET("/generator/handler", generatorHandler.HandlerView())
		web.POST("/generator/handler/execute", generatorHandler.HandlerExecute())

		// 调用方
		web.GET("/authorized/list", authorizedHandler.ListView())
		web.GET("/authorized/add", authorizedHandler.AddView())
		web.GET("/authorized/api/:id", authorizedHandler.ApiView())
		web.GET("/authorized/demo", authorizedHandler.DemoView())

		// 管理员
		web.GET("/admin/list", adminHandler.ListView())
		web.GET("/admin/add", adminHandler.AddView())
		web.GET("/admin/menu", adminHandler.MenuView())
		web.GET("/admin/menu_action/:id", adminHandler.MenuActionView())
		web.GET("/admin/action/:id", adminHandler.AdminMenuView())

		// 升级
		web.GET("/upgrade", upgradeHandler.UpgradeView())
		web.POST("/upgrade/execute", upgradeHandler.UpgradeExecute())

		// 工具箱
		web.GET("/tool/hashids", toolHandler.HashIdsView())
		web.GET("/tool/logs", toolHandler.LogsView())
		web.GET("/tool/cache", toolHandler.CacheView())
		web.GET("/tool/data", toolHandler.DataView())

	}
}

package router

import (
	"github.com/xinliangnote/go-gin-api/internal/web/controller/authorized_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/configinfo_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/dashboard_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/gencode_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/index_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/tool_handler"
)

func setWebRouter(r *resource) {

	indexHandler := index_handler.New(r.logger, r.db, r.cache)
	dashboardHandler := dashboard_handler.New(r.logger, r.db, r.cache)
	genCodeHandler := gencode_handler.New(r.logger, r.db, r.cache)
	configInfoHandler := configinfo_handler.New(r.logger, r.db, r.cache)
	authorizedHandler := authorized_handler.New(r.logger, r.db, r.cache)
	toolHandler := tool_handler.New(r.logger, r.db, r.cache)

	web := r.mux.Group("", r.middles.DisableLog())
	{
		// 首页侧边栏
		web.GET("", indexHandler.View())

		// 仪表盘
		web.GET("/dashboard", dashboardHandler.View())

		// 配置信息
		web.GET("/configinfo", configInfoHandler.View())

		// 代码生成
		web.GET("/init", genCodeHandler.InitView())
		web.POST("/init_exec", genCodeHandler.InitExecute())

		web.GET("/gormgen", genCodeHandler.GormView())
		web.POST("/gormgen_exec", genCodeHandler.GormExecute())

		web.GET("/handlergen", genCodeHandler.HandlerView())
		web.POST("/handlergen_exec", genCodeHandler.HandlerExecute())

		// 调用方
		web.GET("/authorized/list", authorizedHandler.ListView())
		web.GET("/authorized/add", authorizedHandler.AddView())
		web.GET("/authorized/api/:id", authorizedHandler.ApiView())
		web.GET("/authorized/demo", authorizedHandler.DemoView())

		// 工具箱
		web.GET("/tool/hashids", toolHandler.HashIdsView())
		web.GET("/tool/logs", toolHandler.LogsView())

	}
}

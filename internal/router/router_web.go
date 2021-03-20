package router

import (
	"github.com/xinliangnote/go-gin-api/internal/web/controller/configinfo_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/dashboard_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/gencode_handler"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/index_handler"
)

func setWebRouter(r *resource) {

	indexHandler := index_handler.New(r.logger, r.db, r.cache)
	dashboardHandler := dashboard_handler.New(r.logger, r.db, r.cache)
	genCodeHandler := gencode_handler.New(r.logger, r.db, r.cache)
	configInfoHandler := configinfo_handler.New(r.logger, r.db, r.cache)

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

	}
}

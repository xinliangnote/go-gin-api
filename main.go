package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/app/config"
	"go-gin-api/app/route"
	"log"
	"os"
)

func main() {

	//gin.DebugMode   测试模式
	//gin.ReleaseMode 发布模式
	gin.SetMode(gin.DebugMode)
	engine := gin.New()

	// 设置路由
	route.SetupRouter(engine)

	// 启动服务
	if err := engine.Run(config.AppPort); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

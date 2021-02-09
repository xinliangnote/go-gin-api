package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/router"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/pkg/grpc"
	"github.com/xinliangnote/go-gin-api/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/logger"
	"github.com/xinliangnote/go-gin-api/pkg/shutdown"

	"go.uber.org/zap"
)

// @title swagger 接口文档
// @version 2.0
// @description

// @contact.name
// @contact.url
// @contact.email

// @license.name MIT
// @license.url https://github.com/xinliangnote/go-gin-api/blob/master/LICENSE

// @host 127.0.0.1:9999
// @BasePath
func main() {
	// 初始化 logger
	loggers, err := logger.NewJSONLogger(
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName(), env.Active().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.ProjectLogFile()),
	)
	if err != nil {
		panic(err)
	}
	defer loggers.Sync()

	// 初始化 DB
	dbRepo, err := db.New()
	if err != nil {
		loggers.Fatal("new db err", zap.Error(err))
	}

	// 初始化 Cache
	cacheRepo, err := cache.New()
	if err != nil {
		loggers.Fatal("new cache err", zap.Error(err))
	}

	// 初始化 gRPC client
	gRPCRepo, err := grpc.New()
	if err != nil {
		loggers.Fatal("new grpc err", zap.Error(err))
	}

	// 初始化 HTTP 服务
	mux, err := router.NewHTTPMux(loggers, dbRepo, cacheRepo, gRPCRepo)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.ProjectPort(),
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			loggers.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				loggers.Error("server shutdown err", zap.Error(err))
			} else {
				loggers.Info("server shutdown success")
			}
		},

		// 关闭 db
		func() {
			if err := dbRepo.DbWClose(); err != nil {
				loggers.Error("dbw close err", zap.Error(err))
			} else {
				loggers.Info("dbw close success")
			}

			if err := dbRepo.DbRClose(); err != nil {
				loggers.Error("dbr close err", zap.Error(err))
			} else {
				loggers.Info("dbr close success")
			}
		},

		// 关闭 cache
		func() {
			if err := cacheRepo.Close(); err != nil {
				loggers.Error("cache close err", zap.Error(err))
			} else {
				loggers.Info("cache close success")
			}
		},

		// 关闭 gRPC client
		func() {
			if err := gRPCRepo.Conn().Close(); err != nil {
				loggers.Error("gRPC client close err", zap.Error(err))
			} else {
				loggers.Info("gRPC client close success")
			}
		},
	)
}

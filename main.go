package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/configs"
	grpcServer "github.com/xinliangnote/go-gin-api/internal/grpc/server"
	"github.com/xinliangnote/go-gin-api/internal/router"
	"github.com/xinliangnote/go-gin-api/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/logger"
	"github.com/xinliangnote/go-gin-api/pkg/shutdown"
	"github.com/xinliangnote/go-gin-api/pkg/timeutil"

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

// @securityDefinitions.apikey  LoginToken
// @in                          header
// @name                        token

// @BasePath /
func main() {
	// 初始化 access logger
	accessLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(configs.ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}

	// 初始化 cron logger
	cronLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(configs.ProjectCronLogFile),
	)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accessLogger.Sync()
		_ = cronLogger.Sync()
	}()

	// 初始化 HTTP 服务
	s, err := router.NewHTTPServer(accessLogger, cronLogger)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.ProjectPort,
		Handler: s.Mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			accessLogger.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 初始化 gRPC 服务
	gs, err := grpcServer.New(accessLogger, configs.ProjectGRPCPort)
	if err != nil {
		accessLogger.Fatal("grpc server startup err", zap.Error(err))
	}

	go func() {
		if err := gs.Serve(); err != nil {
			accessLogger.Fatal("grpc server serve err", zap.Error(err))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				accessLogger.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if s.Db != nil {
				if err := s.Db.DbWClose(); err != nil {
					accessLogger.Error("dbw close err", zap.Error(err))
				}

				if err := s.Db.DbRClose(); err != nil {
					accessLogger.Error("dbr close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					accessLogger.Error("cache close err", zap.Error(err))
				}
			}
		},

		// 关闭 cron Server
		func() {
			if s.CronServer != nil {
				s.CronServer.Stop()
			}
		},

		// 关闭 gRPC server
		func() {
			if gs != nil {
				gs.GracefulStop()
			}
		},
	)
}

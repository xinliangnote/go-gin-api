package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/router"
	"github.com/xinliangnote/go-gin-api/internal/pkg/configs"
	"github.com/xinliangnote/go-gin-api/pkg/logger"
	"github.com/xinliangnote/go-gin-api/pkg/shutdown"

	"go.uber.org/zap"
)

// @title go-gin-api docs api
// @version
// @description

// @contact.name
// @contact.url
// @contact.email

// @host localhost:9999
// @BasePath
func main() {
	loggers, err := logger.NewJSONLogger(
		logger.WithField("domain", configs.ProjectName()),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(fmt.Sprintf("./logs/%s.log", configs.ProjectName())),
	)
	if err != nil {
		panic(err)
	}

	mux, err := router.NewHTTPMux(loggers)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    ":9999",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			loggers.Fatal("http server startup err", zap.Error(err))
		}
	}()

	shutdown.NewHook().Close(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			loggers.Fatal("shutdown err", zap.Error(err))
		} else {
			loggers.Info("shutdown success")
		}
	})
}

package tool

import (
	"encoding/json"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/pkg/file"

	"go.uber.org/zap"
)

type handler struct {
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
	}
}

func (h *handler) Cache() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("tool_cache", nil)
	}
}

func (h *handler) Data() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("tool_data", nil)
	}
}

func (h *handler) HashIds() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("tool_hashids", configs.Get())
	}
}

func (h *handler) Websocket() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("tool_websocket", nil)
	}
}

func (h *handler) Log() core.HandlerFunc {
	type logData struct {
		Level       string  `json:"level"`
		Time        string  `json:"time"`
		Path        string  `json:"path"`
		HTTPCode    int     `json:"http_code"`
		Method      string  `json:"method"`
		Msg         string  `json:"msg"`
		TraceID     string  `json:"trace_id"`
		Content     string  `json:"content"`
		CostSeconds float64 `json:"cost_seconds"`
	}

	type logsViewResponse struct {
		Logs []logData `json:"logs"`
	}

	type logParseData struct {
		Level        string  `json:"level"`
		Time         string  `json:"time"`
		Caller       string  `json:"caller"`
		Msg          string  `json:"msg"`
		Domain       string  `json:"domain"`
		Method       string  `json:"method"`
		Path         string  `json:"path"`
		HTTPCode     int     `json:"http_code"`
		BusinessCode int     `json:"business_code"`
		Success      bool    `json:"success"`
		CostSeconds  float64 `json:"cost_seconds"`
		TraceID      string  `json:"trace_id"`
	}

	return func(ctx core.Context) {
		readLineFromEnd, err := file.NewReadLineFromEnd(configs.ProjectAccessLogFile)
		if err != nil {
			h.logger.Error("NewReadLineFromEnd err", zap.Error(err))
		}

		logSize := 100

		obj := new(logsViewResponse)
		obj.Logs = make([]logData, logSize)

		for i := 0; i < logSize; i++ {
			content, _ := readLineFromEnd.ReadLine()
			if string(content) != "" {
				var logParse logParseData
				err = json.Unmarshal(content, &logParse)
				if err != nil {
					h.logger.Error("NewReadLineFromEnd json Unmarshal err", zap.Error(err))
				}

				data := logData{
					Content:     string(content),
					Level:       logParse.Level,
					Time:        logParse.Time,
					Path:        logParse.Path,
					Method:      logParse.Method,
					Msg:         logParse.Msg,
					HTTPCode:    logParse.HTTPCode,
					TraceID:     logParse.TraceID,
					CostSeconds: logParse.CostSeconds,
				}

				obj.Logs[i] = data
			}
		}
		ctx.HTML("tool_logs", obj)
	}
}

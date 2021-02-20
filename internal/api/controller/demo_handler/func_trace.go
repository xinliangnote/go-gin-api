package demo_handler

import (
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/third_party_request/go_gin_api_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/httpclient"
	"github.com/xinliangnote/go-gin-api/pkg/p"

	"go.uber.org/zap"
)

type traceResponse []struct {
	Name string `json:"name"` //用户名
	Job  string `json:"job"`  //工作
}

// Trace 示例
// @Summary Trace 示例
// @Description Trace 示例
// @Tags Demo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "签名"
// @Success 200 {object} traceResponse
// @Failure 400 {object} code.Failure
// @Failure 401 {object} code.Failure
// @Router /demo/trace [get]
func (h *handler) Trace() core.HandlerFunc {
	return func(c core.Context) {
		// 三方请求信息
		res1, err := go_gin_api_repo.DemoGet("Tom",
			httpclient.WithTTL(time.Second*5),
			httpclient.WithTrace(c.Trace()),
			httpclient.WithLogger(c.Logger()),
			httpclient.WithHeader("Authorization", c.GetHeader("Authorization")),
			httpclient.WithOnFailedRetry(3, time.Second*1, go_gin_api_repo.DemoGetRetryVerify),
		)

		if err != nil {
			h.logger.Error("get [demo/get] err", zap.Error(err))
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CallHTTPError,
				code.Text(code.CallHTTPError)).WithErr(err),
			)
			return
		}

		// 调试信息
		p.Println("res1.Name", res1.Name, p.WithTrace(c.Trace()))

		// 三方请求信息
		res2, err := go_gin_api_repo.DemoPost("Jack",
			httpclient.WithTTL(time.Second*5),
			httpclient.WithTrace(c.Trace()),
			httpclient.WithLogger(c.Logger()),
			httpclient.WithHeader("Authorization", c.GetHeader("Authorization")),
			httpclient.WithOnFailedRetry(3, time.Second*1, go_gin_api_repo.DemoPostRetryVerify),
		)

		if err != nil {
			h.logger.Error("post [demo/post] err", zap.Error(err))
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CallHTTPError,
				code.Text(code.CallHTTPError)).WithErr(err),
			)
			return
		}

		// 调试信息
		p.Println("res2.Name",
			res2.Name,
			p.WithTrace(c.Trace()),
		)

		// 执行 SQL 信息
		h.userService.GetUserByUserName(c, "test_user")

		// 执行 Redis 信息
		_ = h.cache.Set("name", "tom", time.Minute*10, cache.WithTrace(c.Trace()))
		val, _ := h.cache.Get("name", cache.WithTrace(c.Trace()))
		p.Println("redis-name", val, p.WithTrace(c.Trace()))

		// 初始化客户端
		// client := hello.NewHelloClient(d.grpConn.Conn())
		// client.SayHello(grpc.ContextWithValueAndTimeout(c, time.Second*3), &hello.HelloRequest{Name: "Hello World"})

		data := &traceResponse{
			{
				Name: res1.Name,
				Job:  res1.Job,
			},
			{
				Name: res2.Name,
				Job:  res2.Job,
			},
		}
		c.Payload(data)
	}
}

package demo

import (
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/third_party_request/go_gin_api_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/service/user_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/internal/pkg/grpc"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/httpclient"
	"github.com/xinliangnote/go-gin-api/pkg/p"
	"github.com/xinliangnote/go-gin-api/pkg/token"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Demo struct {
	logger      *zap.Logger
	cache       cache.Repo
	grpconn     grpc.ClientConn
	userService user_service.UserService
}

func NewDemo(logger *zap.Logger, db db.Repo, cache cache.Repo, grpconn grpc.ClientConn) *Demo {
	return &Demo{
		logger:      logger,
		cache:       cache,
		grpconn:     grpconn,
		userService: user_service.NewUserService(db, cache),
	}
}

func (d *Demo) Get() core.HandlerFunc {
	type request struct {
		Name string `uri:"name"`
	}

	type response struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name"`
		Job  string `json:"job"`
	}

	return func(c core.Context) {
		req := new(request)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		if req.Name != "Tom" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.IllegalUserName,
				code.Text(code.IllegalUserName)).WithErr(errors.New("req.Name != Tom")),
			)
			return
		}

		c.Payload(&response{
			Name: "Tom",
			Job:  "Student",
		})
	}
}

func (d *Demo) Post() core.HandlerFunc {
	type request struct {
		Name string `form:"name"`
	}

	type response struct {
		Name string `json:"name"`
		Job  string `json:"job"`
	}

	return func(c core.Context) {
		req := new(request)
		if err := c.ShouldBindPostForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		if req.Name != "Jack" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.IllegalUserName,
				code.Text(code.IllegalUserName)).WithErr(errors.New("req.Name != Jack")),
			)
			return
		}

		c.Payload(&response{
			Name: "Jack",
			Job:  "Teacher",
		})
	}
}

type authResponse struct {
	Authorization string `json:"authorization"` // 签名
	ExpireTime    int64  `json:"expire_time"`   // 过期时间
}

type traceResponse []struct {
	Name string `json:"name"` //用户名
	Job  string `json:"job"`  //工作
}

// 获取授权信息
// @Summary 获取授权信息
// @Description 获取授权信息
// @Tags Demo
// @Accept  json
// @Produce  json
// @Success 200 {object} authResponse
// @Failure 400 {object} code.Failure
// @Router /auth/get [post]
func (d *Demo) Auth() core.HandlerFunc {
	return func(c core.Context) {
		cfg := configs.Get().JWT
		tokenString, err := token.New(cfg.Secret).Sign(1, "xinliangnote", time.Hour*cfg.ExpireDuration)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithErr(err),
			)
			return
		}

		res := new(authResponse)
		res.Authorization = tokenString
		res.ExpireTime = time.Now().Add(time.Hour * cfg.ExpireDuration).Unix()

		c.Payload(res)
	}
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
func (d *Demo) Trace() core.HandlerFunc {
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
			d.logger.Error("get [demo/get] err", zap.Error(err))
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
			d.logger.Error("post [demo/post] err", zap.Error(err))
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
		d.userService.GetUserByUserName(c, "test_user")

		// 执行 Redis 信息
		_ = d.cache.Set("name", "tom", time.Minute*10, cache.WithTrace(c.Trace()))
		val, _ := d.cache.Get("name", cache.WithTrace(c.Trace()))
		p.Println("redis-name", val, p.WithTrace(c.Trace()))

		// 初始化客户端
		// client := hello.NewHelloClient(d.grpconn.Conn())
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

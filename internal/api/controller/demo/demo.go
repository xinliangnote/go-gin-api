package demo

import (
	"net/url"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/jsonparse"
	"github.com/xinliangnote/go-gin-api/pkg/httpclient"

	"go.uber.org/zap"
)

type Demo struct {
	logger *zap.Logger
}

func NewDemo(logger *zap.Logger) *Demo {
	return &Demo{
		logger: logger,
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
			c.SetPayload(code.ErrParam)
			return
		}

		if req.Name != "Tom" {
			c.SetPayload(code.ErrUser)
			return
		}

		c.SetPayload(code.OK.WithData(&response{
			Name: "Tom",
			Job:  "Student",
		}))
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
			c.SetPayload(code.ErrParam)
			return
		}

		if req.Name != "Jack" {
			c.SetPayload(code.ErrUser)
			return
		}

		c.SetPayload(code.OK.WithData(&response{
			Name: "Jack",
			Job:  "Teacher",
		}))
	}
}

type request struct {
	Name string `uri:"name"`
}

type response []struct {
	Name string `json:"name"` //用户名
	Job  string `json:"job"`  //工作
}

// Get 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags Demo
// @Accept  json
// @Produce  json
// @Param name path string true "用户名(Tom)"
// @Param Authorization header string true "签名"
// @Success 200 {object} response "用户信息"
// @Router /demo/user/{name} [get]
func (d *Demo) User() core.HandlerFunc {
	return func(c core.Context) {
		req := new(request)
		if err := c.ShouldBindURI(req); err != nil {
			c.SetPayload(code.ErrParam)
			return
		}

		if req.Name != "Tom" {
			c.SetPayload(code.ErrUser)
			return
		}

		body1, err1 := httpclient.Get("http://127.0.0.1:9999/demo/get/"+req.Name, nil,
			httpclient.WithTTL(time.Second*2),
			httpclient.WithTrace(c.Trace()),
			httpclient.WithLogger(c.Logger()),
			httpclient.WithHeader("Authorization", c.GetHeader("Authorization")),
		)
		if err1 != nil {
			d.logger.Error("get [demo/get] err", zap.Error(err1))
		}

		params := url.Values{}
		params.Set("name", "Jack")
		body2, err2 := httpclient.PostForm("http://127.0.0.1:9999/demo/post", params,
			httpclient.WithTTL(time.Second*2),
			httpclient.WithTrace(c.Trace()),
			httpclient.WithLogger(c.Logger()),
			httpclient.WithHeader("Authorization", c.GetHeader("Authorization")),
		)
		if err2 != nil {
			d.logger.Error("post [demo/post] err", zap.Error(err2))
		}

		data := &response{}
		if err1 == nil && err2 == nil {
			data = &response{
				{
					Name: jsonparse.Get(string(body1), "data.name").(string),
					Job:  jsonparse.Get(string(body1), "data.job").(string),
				},
				{
					Name: jsonparse.Get(string(body2), "data.name").(string),
					Job:  jsonparse.Get(string(body2), "data.job").(string),
				},
			}
		}
		c.SetPayload(code.OK.WithData(data))
	}
}

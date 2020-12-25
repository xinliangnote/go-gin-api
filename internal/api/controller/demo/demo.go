package demo

import (
	"net/url"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/errno"
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
			c.SetPayload(errno.ErrParam)
			return
		}

		if req.Name != "Tom" {
			c.SetPayload(errno.ErrUser)
			return
		}

		c.SetPayload(errno.OK.WithData(&response{
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
			c.SetPayload(errno.ErrParam)
			return
		}

		if req.Name != "Jack" {
			c.SetPayload(errno.ErrUser)
			return
		}

		c.SetPayload(errno.OK.WithData(&response{
			Name: "Jack",
			Job:  "Teacher",
		}))
	}
}

func (d *Demo) User() core.HandlerFunc {

	type response []struct {
		Name string `json:"name"`
	}

	return func(c core.Context) {
		body1, err1 := httpclient.Get("http://127.0.0.1:9999/demo/get/Tom", nil,
			httpclient.WithTTL(time.Second*2),
			httpclient.WithJournal(c.Journal()),
			httpclient.WithLogger(c.Logger()),
		)
		if err1 != nil {
			d.logger.Error("get [demo/get] err", zap.Error(err1))
		}

		params := url.Values{}
		params.Set("name", "Jack")
		body2, err2 := httpclient.PostForm("http://127.0.0.1:9999/demo/post", params,
			httpclient.WithTTL(time.Second*2),
			httpclient.WithJournal(c.Journal()),
			httpclient.WithLogger(c.Logger()),
		)
		if err2 != nil {
			d.logger.Error("post [demo/post] err", zap.Error(err2))
		}

		data := &response{
			{Name: jsonparse.Get(string(body1), "data.name").(string)},
			{Name: jsonparse.Get(string(body2), "data.name").(string)},
		}

		c.SetPayload(errno.OK.WithData(data))
	}
}

package demo

import (
	"fmt"
	"net/url"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/pkg/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/errno"
	"github.com/xinliangnote/go-gin-api/internal/pkg/jsonparse"
	"github.com/xinliangnote/go-gin-api/pkg/aes"
	"github.com/xinliangnote/go-gin-api/pkg/httpclient"
	"github.com/xinliangnote/go-gin-api/pkg/md5"
	"github.com/xinliangnote/go-gin-api/pkg/rsa"

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
			c.SetPayload(errno.ErrParam)
			return
		}

		if req.Name != "Tom" {
			c.SetPayload(errno.ErrUser)
			return
		}

		body1, err1 := httpclient.Get("http://127.0.0.1:9999/demo/get/"+req.Name, nil,
			httpclient.WithTTL(time.Second*2),
			httpclient.WithJournal(c.Journal()),
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
			httpclient.WithJournal(c.Journal()),
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
		c.SetPayload(errno.OK.WithData(data))
	}
}

func (d *Demo) RsaTest() core.HandlerFunc {
	return func(c core.Context) {
		startTime := time.Now()
		encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
		count := 500

		cfg := configs.Get().Rsa
		rsaPublic := rsa.NewPublic(cfg.Public)
		rsaPrivate := rsa.NewPrivate(cfg.Private)

		for i := 0; i < count; i++ {
			// 生成签名
			sn, err := rsaPublic.Encrypt(encryptStr)
			if err != nil {
				d.logger.Error("rsa public encrypt err", zap.Error(err))
			}

			// 验证签名
			_, err = rsaPrivate.Decrypt(sn)
			if err != nil {
				d.logger.Error("rsa private decrypt err", zap.Error(err))
			}
		}
		c.SetPayload(errno.OK.
			WithData(fmt.Sprintf("%v次 - %v", count, time.Since(startTime))),
		)
	}
}

func (d *Demo) AesTest() core.HandlerFunc {
	return func(c core.Context) {
		startTime := time.Now()
		encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
		count := 1000000

		cfg := configs.Get().Aes
		aes := aes.New(cfg.Key, cfg.Iv)
		for i := 0; i < count; i++ {
			// 生成签名
			sn, err := aes.Encrypt(encryptStr)
			if err != nil {
				d.logger.Error("aes encrypt err", zap.Error(err))
			}

			// 验证签名
			_, err = aes.Decrypt(sn)
			if err != nil {
				d.logger.Error("aes decrypt err", zap.Error(err))
			}
		}
		c.SetPayload(errno.OK.
			WithData(fmt.Sprintf("%v次 - %v", count, time.Since(startTime))))
	}
}

func (d *Demo) MD5Test() core.HandlerFunc {
	return func(c core.Context) {
		startTime := time.Now()
		encryptStr := "param_1=xxx&param_2=xxx&ak=xxx&ts=1111111111"
		count := 1000000

		md5 := md5.New()
		for i := 0; i < count; i++ {
			// 生成签名
			md5.Encrypt(encryptStr)

			// 验证签名
			md5.Encrypt(encryptStr)
		}
		c.SetPayload(errno.OK.
			WithData(fmt.Sprintf("%v次 - %v", count, time.Since(startTime))),
		)
	}
}

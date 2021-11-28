package helper

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
	"github.com/xinliangnote/go-gin-api/pkg/signature"
)

type signRequest struct {
	Key    string `form:"key" binding:"required"`    // 调用方 KEY
	Path   string `form:"path" binding:"required"`   // 请求路径 (不附带 querystring)，例如：/api/login
	Method string `form:"method" binding:"required"` // 请求方式，例如：POST
	Params string `form:"params" binding:"required"` // 请求参数，例如：username=tom&password=123456
}

type signResponse struct {
	Authorization     string `json:"authorization"`      // 签名信息-Authorization
	AuthorizationDate string `json:"authorization_date"` // 签名信息-Authorization-Date
}

// Sign 签名
// @Summary 签名
// @Description 签名
// @Tags Helper
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param key formData string true "调用方 KEY"
// @Param path formData string true "请求路径 (不附带 querystring)，例如：/api/login"
// @Param method formData string true "请求方式，例如：POST"
// @Param params formData string true "请求参数，例如：username=tom&password=123456"
// @Success 200 {object} signResponse
// @Failure 400 {object} code.Failure
// @Router /helper/sign [post]
func (h *handler) Sign() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(signRequest)
		res := new(signResponse)

		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		authorizedInfo, err := h.authorizedService.DetailByKey(ctx, req.Key)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(err),
			)
			return
		}

		if authorizedInfo.IsUsed == -1 {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New(req.Key + " 已被禁止调用")),
			)
			return
		}

		fmt.Println(req.Params)

		params, err := url.ParseQuery(req.Params)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				"params 传递格式不正确"),
			)
			return
		}

		sign := signature.New(req.Key, authorizedInfo.Secret, configs.HeaderSignTokenTimeout)
		authorized, date, err := sign.Generate(req.Path, req.Method, params)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				"sign 生成失败"),
			)
			return
		}

		res.Authorization = authorized
		res.AuthorizationDate = date
		ctx.Payload(res)
	}
}

package helper

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type md5Request struct {
	Str string `uri:"str" binding:"required"` // 需要加密的字符串
}

type md5Response struct {
	Md5Str string `json:"md5_str"` // MD5后的字符串
}

// Md5 加密
// @Summary 加密
// @Description 加密
// @Tags Helper
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param str path string true "需要加密的字符串"
// @Success 200 {object} md5Response
// @Failure 400 {object} code.Failure
// @Router /helper/md5/{str} [get]
func (h *handler) Md5() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(md5Request)
		res := new(md5Response)

		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		m := md5.New()
		m.Write([]byte(req.Str))
		res.Md5Str = hex.EncodeToString(m.Sum(nil))
		ctx.Payload(res)
	}
}

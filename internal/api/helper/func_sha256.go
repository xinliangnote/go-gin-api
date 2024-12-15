package helper

import (
    "crypto/sha256"
    "encoding/hex"
    "net/http"

    "github.com/xinliangnote/go-gin-api/internal/code"
    "github.com/xinliangnote/go-gin-api/internal/pkg/core"
    "go.uber.org/zap"
)

type sha256Request struct {
    Str string `uri:"str" binding:"required,max=1024"` // 需要哈希的字符串
}

type sha256Response struct {
    Hash string `json:"hash"` // SHA256 哈希后的字符串
}

// Sha256 生成提供字符串的 SHA256 哈希。
// @Summary 生成 SHA256 哈希
// @Description 生成提供字符串的 SHA256 哈希
// @Tags Helper
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param str path string true "需要使用 SHA256 哈希的字符串"
// @Success 200 {object} sha256Response "成功生成哈希字符串"
// @Failure 400 {object} code.Failure "输入参数无效"
// @Failure 500 {object} code.Failure "内部服务器错误"
// @Router /helper/sha256/{str} [get]
func (h *handler) Sha256() core.HandlerFunc {
    return func(ctx core.Context) {
        req := new(sha256Request)
        res := new(sha256Response)

        if err := ctx.ShouldBindURI(req); err != nil {
            h.logger.Error("绑定 URI 参数失败", zap.Error(err))
            ctx.AbortWithError(core.Error(
                http.StatusBadRequest,
                code.ParamBindError,
                code.Text(code.ParamBindError)).WithError(err),
            )
            return
        }

        hash := sha256.Sum256([]byte(req.Str))
        res.Hash = hex.EncodeToString(hash[:])
        h.logger.Info("成功生成 SHA256 哈希", zap.String("输入", req.Str), zap.String("哈希", res.Hash))
        ctx.Payload(res)
    }
}
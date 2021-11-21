package tool

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type clearCacheRequest struct {
	RedisKey string `form:"redis_key"` // Redis Key
}

type clearCacheResponse struct {
	Bool bool `json:"bool"` // 删除结果
}

// ClearCache 清空缓存
// @Summary 清空缓存
// @Description 清空缓存
// @Tags API.tool
// @Accept multipart/form-data
// @Produce json
// @Param redis_key formData string true "Redis Key"
// @Success 200 {object} searchCacheResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/cache/clear [patch]
func (h *handler) ClearCache() core.HandlerFunc {
	return func(c core.Context) {
		req := new(clearCacheRequest)
		res := new(clearCacheResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		if b := h.cache.Exists(req.RedisKey); b != true {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CacheNotExist,
				code.Text(code.CacheNotExist)),
			)
			return
		}

		b := h.cache.Del(req.RedisKey, redis.WithTrace(c.Trace()))
		if b != true {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.CacheDelError,
				code.Text(code.CacheDelError)),
			)
			return
		}

		res.Bool = b
		c.Payload(res)
	}
}

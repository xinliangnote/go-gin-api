## 执行命令

```$xslt
// test_handler 为 ./internal/api/controller/ 中的包名
./scripts/handlergen.sh test_handler
```

## 模板文件参考

```go
package test_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/api/service/user_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	// i 为了避免被其他包实现
	i()

	// Create 创建用户
	// @Tags Test
	// @Router /test/create [post]
	Create() core.HandlerFunc

	// Update 编辑用户
	// @Tags Test
	// @Router /test/update [post]
	Update() core.HandlerFunc

	// Delete 删除用户
	// @Tags Test
	// @Router /test/delete [post]
	Delete() core.HandlerFunc

	// Detail 用户详情
	// @Tags Test
	// @Router /test/detail [post]
	Detail() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       cache.Repo
	userService user_service.UserService
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		userService: user_service.NewUserService(db, cache),
	}
}

func (h *handler) i() {}

```

以上会生成 4 个文件
- func_create.go
- func_update.go
- func_delete.go
- func_detail.go

## func_create.go 参考

```go
package test_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type createRequest struct{}

type createResponse struct{}

// Create 创建用户
// @Summary 创建用户
// @Description 创建用户
// @Tags Test
// @Accept json
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /test/create [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {

	}
}

```
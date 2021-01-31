package resolvers

import (
	"context"

	"github.com/xinliangnote/go-gin-api/internal/api/service/user_service"
	"github.com/xinliangnote/go-gin-api/internal/graph/generated"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"

	"go.uber.org/zap"
)

type coreCtxKeyType struct{ name string }

var CoreContextKey = coreCtxKeyType{"_core_context"}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

type Resolver struct {
	logger      *zap.Logger
	cache       cache.Repo
	userService user_service.UserService
}

func NewRootResolvers(logger *zap.Logger, db db.Repo, cache cache.Repo) generated.Config {
	c := generated.Config{
		Resolvers: &Resolver{
			logger:      logger,
			cache:       cache,
			userService: user_service.NewUserService(db, cache),
		},
	}
	return c
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

// getCoreContextByCtx 获取 core context
func (r *Resolver) getCoreContextByCtx(ctx context.Context) core.Context {
	return ctx.Value(CoreContextKey).(core.Context)
}

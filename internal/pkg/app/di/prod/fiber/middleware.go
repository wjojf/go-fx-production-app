package fiber

import (
	"github.com/wjojf/go-uber-fx/internal/api/http/middleware"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"go.uber.org/fx"
)

func MiddlewareStack(cfg config.Config) fx.Option {
	return fx.Provide(
		fx.Annotate(
			middleware.DummyMiddleware,
			fx.ResultTags(`group:"middlewares"`),
		),
	)
}

package fiber

import (
	"github.com/wjojf/go-uber-fx/internal/api/http/middleware"
	"go.uber.org/fx"
)

var MiddlewareStack = fx.Provide(
	fx.Annotate(
		middleware.DummyMiddleware,
		fx.ResultTags(`group:"middlewares"`),
	),
)

package fiber

import (
	"github.com/gofiber/fiber/v3/middleware/pprof"
	"github.com/wjojf/go-uber-fx/internal/api/http/middleware"
	"go.uber.org/fx"
)

var MiddlewareStack = fx.Provide(
	// // Dummy, remove later
	// fx.Annotate(
	// 	middleware.DummyMiddleware,
	// 	fx.ResultTags(`group:"middlewares"`),
	// ),

	// Request Logger
	fx.Annotate(
		middleware.RequestLogger,
		fx.ResultTags(`group:"middlewares"`),
	),

	// Profiling
	fx.Annotate(
		pprof.New,
		fx.ResultTags(`group:"middlewares"`),
	),
)

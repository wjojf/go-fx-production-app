package fiber

import (
	"github.com/wjojf/go-uber-fx/internal/api/http"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"fiber-prod",

		// Middleware
		MiddlewareStack,

		// Server Options
		Options,

		// Server Default Bundle
		http.Module,
	)
}

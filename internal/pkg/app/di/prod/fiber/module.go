package fiber

import (
	"github.com/wjojf/go-uber-fx/internal/api/http"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"go.uber.org/fx"
)

func Module(cfg config.Config) fx.Option {
	return fx.Module(
		"fiber-prod",

		// Server Config
		fx.Supply(ConfigProd),

		// Server Middleware
		MiddlewareStack(cfg),

		// Server Default Bundle
		http.Module,
	)
}

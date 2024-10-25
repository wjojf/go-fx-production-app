package fiber

import (
	"github.com/wjojf/go-uber-fx/internal/api/http"
	"github.com/wjojf/go-uber-fx/internal/pkg/app/di/prod/graphql"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"fiber-prod",

		// Middleware
		MiddlewareStack,

		// Server Options
		Options,

		// GraphQL Routes
		fx.Invoke(
			graphql.AddGraphQLRoutes,
		),

		// Server Default Bundle
		http.Module,
	)
}

package graphql

import (
	"github.com/wjojf/go-uber-fx/internal/api/graphql/resolvers"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"graphql-api",

	// Resolvers
	fx.Provide(
		resolvers.New,
	),

	// Server
	fx.Provide(
		NewServer,
	),
)
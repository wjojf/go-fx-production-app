package http

import (
	"github.com/wjojf/go-uber-fx/internal/api/http/handlers"
	"github.com/wjojf/go-uber-fx/internal/api/http/routes"
	"github.com/wjojf/go-uber-fx/internal/api/http/server"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http-api",

	// Register *fiber.App
	fx.Provide(
		fx.Annotate(
			server.New,
			fx.ParamTags(`group:"middlewares"`),
		),
	),

	// Register All Handlers
	handlers.Module,

	// Register Routes
	routes.Module,
)

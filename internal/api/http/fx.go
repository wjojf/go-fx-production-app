package http

import (
	"github.com/wjojf/go-uber-fx/internal/api/http/handlers"
	"github.com/wjojf/go-uber-fx/internal/api/http/routes"
	"github.com/wjojf/go-uber-fx/internal/api/http/server"
	"github.com/wjojf/go-uber-fx/internal/api/http/service"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http-api",

	// Register *fiber.App
	fx.Provide(
		fx.Annotate(
			server.New,
			fx.ParamTags(`group:"options"`),
		),
	),

	// Register All Handlers
	handlers.Module,

	// Register All Services
	service.Module,

	// Register Routes
	routes.Module,
)

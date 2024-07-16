package http

import (
	"github.com/wjojf/go-uber-fx/internal/api/http/handlers"
	"github.com/wjojf/go-uber-fx/internal/api/http/routes"
	"github.com/wjojf/go-uber-fx/internal/api/http/server"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http-server",

	// Register *fiber.App
	fx.Provide(server.New),

	// Register All Handlers
	handlers.Module,

	// Register Routes
	routes.Module,
)

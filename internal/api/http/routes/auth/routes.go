package auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/wjojf/go-uber-fx/internal/api/http/handlers/auth"
)

func AddRoutes(app *fiber.App, handler auth.Handler) {

	group := app.Group("/api/v1/auth")

	group.Post("/login", handler.Login)
	group.Post("/refresh", handler.Refresh)

}

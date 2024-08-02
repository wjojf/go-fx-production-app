package users

import (
	"github.com/gofiber/fiber/v3"
	handler "github.com/wjojf/go-uber-fx/internal/api/http/handlers/users"
)

func AddRoutes(app *fiber.App, h handler.Handler) {
	group := app.Group("/api/v1/users")

	group.Get("/", h.GetAll)
	group.Get("/:id", h.GetByID)
	group.Post("/", h.CreateUser)
	group.Put("/:id", h.UpdateUserFull)
	group.Patch("/:id", h.UpdateUserPartial)
}

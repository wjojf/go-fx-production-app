package users

import (
	"github.com/gofiber/fiber/v3"
	handler "github.com/wjojf/go-uber-fx/internal/api/http/handlers/users"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
)

func AddRoutes(app *fiber.App, cfg config.Config, h handler.Handler) {
	group := app.Group("/api/v1/users")
	// .Use(middleware.CheckAuthentication(cfg, h.Logger(), h.Repository(), false))

	group.Get("/", h.GetAll)
	group.Get("/:id", h.GetByID)
	group.Post("/", h.CreateUser)
	group.Put("/:id", h.UpdateUserFull)
	group.Patch("/:id", h.UpdateUserPartial)
}

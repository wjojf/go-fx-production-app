package index

import (
	"github.com/gofiber/fiber/v3"
	handler "github.com/wjojf/go-uber-fx/internal/api/http/handlers/index"
)

func AddRoute(app *fiber.App, h handler.Handler) {
	app.Get("/", h.Index)
}

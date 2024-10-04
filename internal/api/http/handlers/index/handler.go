package index

import (
	"github.com/gofiber/fiber/v3"
)

type Handler struct{}

func New() Handler {
	return Handler{}
}

func (h Handler) Index(ctx fiber.Ctx) error {
	return ctx.Render("static/html/index.html", fiber.Map{
		"Ip": ctx.IP(),
	})
}

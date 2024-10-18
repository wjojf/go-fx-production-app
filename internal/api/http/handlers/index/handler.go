package index

import (
	"github.com/gofiber/fiber/v3"
)

const (
	indexTemplate = "static/html/index.html"
)

type Handler struct{}

func New() Handler {
	return Handler{}
}

func (h Handler) Index(ctx fiber.Ctx) error {
	return ctx.Render(indexTemplate, fiber.Map{"Ip": ctx.IP()})
}

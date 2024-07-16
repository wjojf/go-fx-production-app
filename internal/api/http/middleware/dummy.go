package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

func DummyMiddleware(log *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		log.Debug("I'm a dummy middleware!")
		return c.Next()
	}
}

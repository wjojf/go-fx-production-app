package middleware

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
)

func RequestLogger(logger *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		startTime := time.Now()

		// Process the request
		err := c.Next()

		duration := time.Since(startTime).Milliseconds()

		// Log request information
		logger.Info("Request Processed",
			slog.String("request_id", string(c.Request().Header.Peek("X-Request-ID"))),
			slog.String("method", c.Method()),
			slog.String("url", c.OriginalURL()),
			slog.Int("status", c.Response().StatusCode()),
			slog.Int64("latency_ms", duration),
			slog.String("client_ip", c.IP()),
		)

		return err
	}
}

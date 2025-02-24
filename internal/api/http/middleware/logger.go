package middleware

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/wjojf/go-uber-fx/internal/api/http/service"
	"github.com/wjojf/go-uber-fx/internal/api/http/utils"
)

func RequestLogger(logger *slog.Logger, jwtService service.JwtService) fiber.Handler {
	return func(c fiber.Ctx) error {
		startTime := time.Now()

		// Process the request
		err := c.Next()

		duration := time.Since(startTime).Milliseconds()

		token, _ := utils.ExtractAuthToken(c.Get("Authorization"))
		userId, _ := jwtService.ExtractUserID(token)

		// Log request information
		logger.Info("HTTP Request",
			slog.String("request_id", string(c.Request().Header.Peek("X-Request-ID"))),
			slog.String("method", c.Method()),
			slog.String("url", c.OriginalURL()),
			slog.Int("status", c.Response().StatusCode()),
			slog.Int64("latency_ms", duration),
			slog.String("client_ip", c.IP()),
			slog.String("user_id", userId),
		)

		return err
	}
}

package middleware

import (
	"fmt"
	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
	types "github.com/wjojf/go-uber-fx/internal/api/http/types/auth"
	"github.com/wjojf/go-uber-fx/internal/api/http/utils"
	"github.com/wjojf/go-uber-fx/internal/domain/users/repository"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"github.com/wjojf/go-uber-fx/pkg/jwt"
	"log/slog"
	"time"
)

func CheckAuthentication(cfg config.Config, log *slog.Logger, r repository.UsersRepository, checkStorage bool) fiber.Handler {
	return func(ctx fiber.Ctx) error {

		token, err := utils.ExtractAuthToken(ctx.Get("Authorization"))

		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": errors.Wrap(err, "invalid authorization format").Error(),
			})
		}

		if token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized. No Authorization header",
			})
		}

		// Parse token payload
		var claims = &types.JwtPayload{}
		if err := jwt.DecodeToken(token, jwtLib.SigningMethodHS256, cfg.JwtSigningKey, claims); err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": fmt.Sprintf("Unauthorized. Error decoding token: %v", err.Error()),
			})
		}

		// Check if token expired
		if claims.ExpiresAt < time.Now().Unix() {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized. Token expired",
			})
		}

		// Check if user exists
		if checkStorage {
			_, err := r.GetUserByID(ctx.Context(), claims.UserId)
			if err != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized. User not found",
				})
			}
		}

		ctx.Set(utils.UserIdContextKey, claims.UserId)

		return ctx.Next()
	}
}

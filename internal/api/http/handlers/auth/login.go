package auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
	"github.com/wjojf/go-uber-fx/internal/api/http/types/auth"
)

func (h Handler) Login(ctx fiber.Ctx) error {

	var dto auth.LoginRequestDTO
	if err := ctx.Bind().JSON(&dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.r.GetUserByEmail(ctx.Context(), dto.Email)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": errors.Wrap(err, "failed to get user by email"),
		})
	}

	tokens, err := h.jwtService.GenerateTokens(user.ID())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errors.Wrap(err, "failed to generate tokens"),
		})
	}

	return ctx.JSON(auth.LoginResponseDTO{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})

}

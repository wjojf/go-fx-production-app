package auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
	types "github.com/wjojf/go-uber-fx/internal/api/http/types/auth"
)

func (h Handler) Refresh(ctx fiber.Ctx) error {

	var dto types.RefreshRequestDTO
	if err := ctx.Bind().JSON(&dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors.Wrap(err, "invalid incoming body").Error(),
		})
	}

	refreshedTokenData, err := h.jwtService.RefreshAccessToken(dto.RefreshToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": errors.Wrap(err, "invalid refresh token").Error(),
		})
	}

	// Check if user still exists
	if _, err := h.r.GetUserByID(ctx.Context(), refreshedTokenData.UserID); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": errors.Wrap(err, "user not found").Error(),
		})
	}

	return ctx.JSON(types.RefreshResponseDTO{AccessToken: refreshedTokenData.AccessToken})
}

package users

import (
	"github.com/gofiber/fiber/v3"
	schemas "github.com/wjojf/go-uber-fx/internal/api/http/types/users"
)

func (h Handler) GetByID(c fiber.Ctx) error {

	userID := c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is required"})
	}

	user, err := h.r.GetUserByID(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(schemas.FromUser(user))
}

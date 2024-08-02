package users

import (
	"github.com/gofiber/fiber/v3"
	schemas "github.com/wjojf/go-uber-fx/internal/api/http/types/users"
)

func (h Handler) GetAll(c fiber.Ctx) error {
	users, err := h.r.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var response []schemas.UserResponse
	for _, user := range users {
		response = append(response, schemas.FromUser(user))
	}

	return c.Status(200).JSON(response)
}

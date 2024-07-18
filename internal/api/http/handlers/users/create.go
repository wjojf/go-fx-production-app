package users

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	schemas "github.com/wjojf/go-uber-fx/internal/api/http/types/users"
)

func (h Handler) CreateUser(c fiber.Ctx) error {

	var createUserRequest schemas.CreateUserRequest
	if err := c.Bind().JSON(&createUserRequest); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": fmt.Sprintf("Cannot parse JSON: %s", err.Error()),
			},
		)
	}

	vo, err := createUserRequest.ToValueObject()
	if err != nil {
		return c.Status(422).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	user, err := h.r.SaveUser(vo)
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(schemas.FromUser(user))
}

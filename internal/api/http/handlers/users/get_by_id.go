package users

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func (h Handler) GetByID(c fiber.Ctx) error {

	user, err := h.r.GetUserByID(1)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	h.log.Info(fmt.Sprintf("User: %+v", user))

	return c.JSON(
		fiber.Map{
			"message": "User found!",
		},
	)
}

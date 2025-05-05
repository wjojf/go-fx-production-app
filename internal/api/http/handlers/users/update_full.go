package users

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	schemas "github.com/wjojf/go-uber-fx/internal/api/http/types/users"
)

func (h Handler) UpdateUserFull(c fiber.Ctx) error {

	var userID = c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "User ID is required",
			},
		)
	}

	var updateFullRequest schemas.UpdateFullRequest
	if err := c.Bind().JSON(&updateFullRequest); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": fmt.Sprintf("Cannot parse JSON: %s", err.Error()),
			},
		)
	}

	vo, err := updateFullRequest.ToValueObject()
	if err != nil {
		return c.Status(422).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	user, err := h.r.UpdateUserByID(c.UserContext(), userID, vo)
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(schemas.FromUser(user))

}

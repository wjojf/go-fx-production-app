package users

import (
	"fmt"
	"github.com/opentracing/opentracing-go"

	"github.com/gofiber/fiber/v3"
	schemas "github.com/wjojf/go-uber-fx/internal/api/http/types/users"
)

func (h Handler) UpdateUserPartial(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "UpdateUserPartial")
	defer span.Finish()

	var userID = c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "User ID is required",
			},
		)
	}

	var updatePartialRequest schemas.UpdatePartialRequest
	if err := c.Bind().JSON(&updatePartialRequest); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": fmt.Sprintf("Cannot parse JSON: %s", err.Error()),
			},
		)
	}

	vo, err := updatePartialRequest.ToValueObject()
	if err != nil {
		return c.Status(422).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	user, err := h.r.UpdateUserByID(ctx, userID, vo)
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(schemas.FromUser(user))
}

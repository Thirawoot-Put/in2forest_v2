package handlers

import (
	"thirawoot/in2forest_shop_backend/internal/application"

	"github.com/gofiber/fiber/v2"
)

func HandleAppErr(e error, c *fiber.Ctx) error {
	if appErr, ok := e.(*application.AppErr); ok {
		return c.Status(appErr.Code).JSON(fiber.Map{
			"message": appErr.Message,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Unexpected error",
	})
}

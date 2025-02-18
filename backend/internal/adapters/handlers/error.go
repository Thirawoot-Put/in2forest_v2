package handlers

import (
	"log"
	"thirawoot/in2forest_shop_backend/internal/application"

	"github.com/gofiber/fiber/v2"
)

func HandleAppErr(e error, c *fiber.Ctx) error {
	log.Printf("Error occurred in %s %s: %v \n", c.Method(), c.Path(), e)

	if appErr, ok := e.(*application.AppErr); ok {
		return c.Status(appErr.Code).JSON(fiber.Map{
			"message": appErr.Message,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Unexpected error",
	})
}

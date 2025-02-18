package handlers

import (
	"log"
	"thirawoot/in2forest_shop_backend/internal/application"

	"github.com/gofiber/fiber/v2"
)

func HandleAppErr(e error, c *fiber.Ctx) error {
	if appErr, ok := e.(*application.AppErr); ok {
		log.Printf("Error in %s %s: %s \n", c.Method(), c.Path(), appErr.Error())
		return c.Status(appErr.Code).JSON(fiber.Map{
			"message": appErr.Message,
		})
	}

	log.Printf("Error in %s %s: %v \n", c.Method(), c.Path(), e)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Unexpected error",
	})
}

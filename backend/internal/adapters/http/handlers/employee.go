package handlers

import (
	"fmt"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type EmployeeHandler struct {
	app portin.EmployeeApp
}

func NewEmployeeHandler(app portin.EmployeeApp) *EmployeeHandler {
	return &EmployeeHandler{app: app}
}

func (h *EmployeeHandler) GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	fmt.Println("local ---> ", claims)
	return c.SendString("profile")
}

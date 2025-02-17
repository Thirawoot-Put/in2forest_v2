package handlers

import (
	"fmt"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"

	"github.com/gofiber/fiber/v2"
)

type EmployeeRoleHandler struct {
	service portin.EmployeeRoleService
}

func NewEmployeeRoleHandler(service portin.EmployeeRoleService) EmployeeRoleHandler {
	return EmployeeRoleHandler{service: service}
}

func (s *EmployeeRoleHandler) PostEmployeeRole(c *fiber.Ctx) error {
	fmt.Printf("req body --> %v", c.Body())
	fmt.Printf("req body in request --> %v", c.Request().Body())
	c.Response().BodyWriter().Write([]byte("Hello, World!"))
	// => "Hello, World!"
	return nil
}

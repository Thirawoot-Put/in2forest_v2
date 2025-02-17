package handlers

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
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
	var input dto.EmployeeRoleCreate
	err := c.BodyParser(&input)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "failed to parse request body")
	}

	result := s.service.Create(input)

	return c.Status(result.StatusCode).JSON(result)
}

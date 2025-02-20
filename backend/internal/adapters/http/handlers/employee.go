package handlers

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"
	"thirawoot/in2forest_shop_backend/pkg"

	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	app portin.EmployeeApp
}

func NewEmployeeHandler(app portin.EmployeeApp) *EmployeeHandler {
	return &EmployeeHandler{app: app}
}

func (h *EmployeeHandler) GetEmployeeProfile(c *fiber.Ctx) error {
	user := c.Locals("user")
	id := pkg.GetSubInToken(user)
	if id == 0 {
		c.Status(constants.Code.Forbidden).JSON(ApiResponse("error", "JWT_MALFORMED"))
	}

	result, err := h.app.Find(uint(id))
	if err != nil {
		return appErrHandler(err, c)
	}

	return c.Status(constants.Code.Ok).JSON(ApiResponse("success", result))
}

func (h *EmployeeHandler) PutEmployee(c *fiber.Ctx) error {
	var input dto.Employee
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(constants.Code.BadRequest).JSON(ApiResponse("error", "FAILED_TO_PARSE_REQUEST_BODY"))
	}

	user := c.Locals("user")
	id := pkg.GetSubInToken(user)
	if id == 0 {
		return c.Status(constants.Code.Forbidden).JSON(ApiResponse("error", "JWT_MALFORMED"))
	}

	result, err := h.app.Update(id, input)
	if err != nil {
		return appErrHandler(err, c)
	}

	return c.Status(constants.Code.Ok).JSON(ApiResponse("success", result))
}

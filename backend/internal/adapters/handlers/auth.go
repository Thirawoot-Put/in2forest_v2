package handlers

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"

	"github.com/gofiber/fiber/v2"
)

type AuthEmployeeHandler struct {
	app portin.AuthEmployeeApp
}

func NewAuthEmployeeHandler(app portin.AuthEmployeeApp) *AuthEmployeeHandler {
	return &AuthEmployeeHandler{
		app: app,
	}
}

func (h *AuthEmployeeHandler) RegisterAdmin(c *fiber.Ctx) error {
	var input dto.Employee
	err := c.BodyParser(&input)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "FAILED_TO_PARSE_REQUEST_BODY")
	}

	result, err := h.app.RegisterAdmin(input)
	if err != nil {
		return HandleAppErr(err, c)
	}

	return c.Status(constants.Code.Created).JSON(ApiResponse(result))
}

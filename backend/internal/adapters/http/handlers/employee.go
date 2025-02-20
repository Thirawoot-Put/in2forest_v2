package handlers

import (
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

func (h *EmployeeHandler) GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user")
	id := pkg.GetSubInToken(user)

	result, err := h.app.Find(uint(id))
	if err != nil {
		return appErrHandler(err, c)
	}

	return c.Status(constants.Code.Ok).JSON(ApiResponse("success", result))
}

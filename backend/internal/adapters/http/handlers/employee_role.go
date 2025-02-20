package handlers

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"

	"github.com/gofiber/fiber/v2"
)

type EmployeeRoleHandler struct {
	app portin.EmployeeRoleApp
}

func NewEmployeeRoleHandler(app portin.EmployeeRoleApp) *EmployeeRoleHandler {
	return &EmployeeRoleHandler{app: app}
}

func (h *EmployeeRoleHandler) PostEmployeeRole(c *fiber.Ctx) error {
	var input dto.EmployeeRole
	err := c.BodyParser(&input)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "FAILED_TO_PARSE_REQUEST_BODY")
	}

	result, err := h.app.Create(input)
	if err != nil {
		return appErrHandler(err, c)
	}

	return c.Status(constants.Code.Created).JSON(ApiResponse("success", result))
}

func (h *EmployeeRoleHandler) DeleteEmployeeRole(c *fiber.Ctx) error {
	id, err := atoiParam(c.AllParams()["id"])
	if err != nil {
		return fiber.NewError(constants.Code.BadRequest, err.Error())
	}

	result, err := h.app.Delete(uint(id))
	if err != nil {
		return appErrHandler(err, c)
	}

	return c.Status(constants.Code.Ok).JSON(ApiResponse("success", result))
}

func (h *EmployeeRoleHandler) PutEmployeeRole(c *fiber.Ctx) error {
	id, err := atoiParam(c.AllParams()["id"])
	if err != nil {
		return fiber.NewError(constants.Code.BadRequest, err.Error())
	}

	var input dto.EmployeeRole
	err = c.BodyParser(&input)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "FAILED_TO_PARSE_REQUEST_BODY")
	}

	result, err := h.app.Update(uint(id), input)
	if err != nil {
		return appErrHandler(err, c)
	}

	return c.Status(constants.Code.Ok).JSON(ApiResponse("success", result))
}

func (h *EmployeeRoleHandler) GetEmployeeRole(c *fiber.Ctx) error {
	id, err := atoiParam(c.AllParams()["id"])
	if err != nil {
		return fiber.NewError(constants.Code.BadRequest, err.Error())
	}

	result, err := h.app.Find(uint(id))
	if err != nil {
		return appErrHandler(err, c)
	}

	return c.Status(constants.Code.Ok).JSON(ApiResponse("success", result))
}

func (h *EmployeeRoleHandler) GetEmployeeRoles(c *fiber.Ctx) error {
	result, err := h.app.FindAll()
	if err != nil {
		return appErrHandler(err, c)
	}

	return c.Status(constants.Code.Ok).JSON(ApiResponse("success", result))
}

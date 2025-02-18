package handlers

import (
	"fmt"
	"strconv"
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"

	"github.com/gofiber/fiber/v2"
)

type EmployeeRoleHandler struct {
	service portin.EmployeeRoleApp
}

func NewEmployeeRoleHandler(service portin.EmployeeRoleApp) EmployeeRoleHandler {
	return EmployeeRoleHandler{service: service}
}

func atoiParam(param string) (int, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return -1, fmt.Errorf("FAILED_TO_CONVERT_ID_TO_INT")
	}

	return id, nil
}

func (h *EmployeeRoleHandler) PostEmployeeRole(c *fiber.Ctx) error {
	var input dto.EmployeeRoleCreate
	err := c.BodyParser(&input)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "FAILED_TO_PARSE_REQUEST_BODY")
	}

	result, err := h.service.Create(input)
	if err != nil {
		return HandleAppErr(err, c)
	}

	return c.Status(constants.Code.Created).JSON(ApiResponse(result))
}

func (h *EmployeeRoleHandler) DeleteEmployeeRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil {
		return fiber.NewError(constants.Code.BadRequest, "FAILED_TO_CONVERT_ID_TO_INT")
	}

	result, err := h.service.Delete(uint(id))
	if err != nil {
		return HandleAppErr(err, c)
	}

	return c.Status(constants.Code.Ok).JSON(ApiResponse(result))
}

func (h *EmployeeRoleHandler) PutEmployeeRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.AllParams()["id"])
	if err != nil {
		return fiber.NewError(constants.Code.BadRequest, "FAILED_TO_CONVERT_ID_TO_INT")
	}

	var input dto.EmployeeRoleCreate
	err = c.BodyParser(&input)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "FAILED_TO_PARSE_REQUEST_BODY")
	}

	result, err := h.service.Update(uint(id), input)
	if err != nil {
		return HandleAppErr(err, c)
	}

	return c.Status(constants.Code.Ok).JSON(ApiResponse(result))
}

func (h *EmployeeRoleHandler) GetEmployeeRole(c *fiber.Ctx) error {
	id, err := atoiParam(c.AllParams()["id"])
	if err != nil {
		return fiber.NewError(constants.Code.BadRequest, err.Error())
	}

	result, err := h.service.Find(uint(id))
	if err != nil {
		return HandleAppErr(err, c)
	}

	return c.Status(constants.Code.Ok).JSON(ApiResponse(result))
}

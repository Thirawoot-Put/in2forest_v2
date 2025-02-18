package constants

import "github.com/gofiber/fiber/v2"

type ApiStatus struct {
	Ok           int
	Created      int
	BadRequest   int
	Unauthorized int
	Forbidden    int
	NotFound     int
	Conflict     int
	ServerError  int
}

var Code = ApiStatus{
	Ok:           fiber.StatusOK,
	Created:      fiber.StatusCreated,
	BadRequest:   fiber.StatusBadRequest,
	Unauthorized: fiber.StatusUnauthorized,
	Forbidden:    fiber.StatusForbidden,
	NotFound:     fiber.StatusNotFound,
	Conflict:     fiber.StatusConflict,
	ServerError:  fiber.StatusInternalServerError,
}

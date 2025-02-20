package handlers

import "github.com/gofiber/fiber/v2"

// func ApiResponse(data interface{}) map[string]interface{} {
// 	return map[string]interface{}{"message": "success", "data": data}
// }

func ApiResponse(message string, data interface{}) fiber.Map {
	if message == "" {
		message = "success"
	}

	return fiber.Map{"message": message, "data": data}
}

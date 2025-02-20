package middlewares

import (
	// "thirawoot/in2forest_shop_backend/internal/adapters/http/handlers"
	// "thirawoot/in2forest_shop_backend/internal/utils/constants"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v5"
)

func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
	})
}

// func RoleMiddleware(requiredRole string) fiber.Handler {
//   return func(c *fiber.Ctx) error {
//     user := c.Locals("user").(*jwt.Token)
//     if user == nil {
//       return c.Status(constants.Code.Unauthorized).JSON(handlers.ApiResponse())
//     }
//   }
// }

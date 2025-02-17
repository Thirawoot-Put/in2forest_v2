package routes

import (
	"thirawoot/in2forest_shop_backend/internal/adapters/handlers"
	"thirawoot/in2forest_shop_backend/internal/application"
	"thirawoot/in2forest_shop_backend/internal/infras/database/repositories"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func EmployeeRoleRoutes(route fiber.Router, db *gorm.DB) {
	repository := repositories.NewEmployeeRoleRepository(db)
	service := application.NewEmployeeRoleService(repository)
	handlers := handlers.NewEmployeeRoleHandler(service)

	employeeRoleRoute := route.Group("/employee-role")
	{
		employeeRoleRoute.Post("", handlers.PostEmployeeRole)
		employeeRoleRoute.Delete("/:id", handlers.Delete)
	}
}

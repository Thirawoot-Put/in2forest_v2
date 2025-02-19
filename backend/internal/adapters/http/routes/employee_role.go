package routes

import (
	"thirawoot/in2forest_shop_backend/internal/adapters/database/repositories"
	"thirawoot/in2forest_shop_backend/internal/adapters/http/handlers"
	"thirawoot/in2forest_shop_backend/internal/application"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func EmployeeRoleRoutes(route *fiber.Router, db *gorm.DB) {
	r := *route

	repository := repositories.NewEmployeeRoleRepository(db)
	service := application.NewEmployeeRoleApp(repository)
	handlers := handlers.NewEmployeeRoleHandler(service)

	empyRoleRoute := r.Group("/employee-role")
	{
		empyRoleRoute.Post("", handlers.PostEmployeeRole)
		empyRoleRoute.Get("/", handlers.GetEmployeeRoles)
		empyRoleRoute.Delete("/:id", handlers.DeleteEmployeeRole)
		empyRoleRoute.Put("/:id", handlers.PutEmployeeRole)
		empyRoleRoute.Get("/:id", handlers.GetEmployeeRole)
	}
}

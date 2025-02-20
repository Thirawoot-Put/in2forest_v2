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

	repo := repositories.NewEmployeeRoleRepository(db)
	app := application.NewEmployeeRoleApp(repo)
	handler := handlers.NewEmployeeRoleHandler(app)

	empRoleRoute := r.Group("/employee-role")
	{
		empRoleRoute.Post("", handler.PostEmployeeRole)
		empRoleRoute.Get("/", handler.GetEmployeeRoles)
		empRoleRoute.Delete("/:id", handler.DeleteEmployeeRole)
		empRoleRoute.Put("/:id", handler.PutEmployeeRole)
		empRoleRoute.Get("/:id", handler.GetEmployeeRole)
	}
}

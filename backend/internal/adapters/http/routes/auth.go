package routes

import (
	"thirawoot/in2forest_shop_backend/internal/adapters/database/repositories"
	"thirawoot/in2forest_shop_backend/internal/adapters/http/handlers"
	"thirawoot/in2forest_shop_backend/internal/application"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoutes(route *fiber.Router, db *gorm.DB) {
	r := *route

	authRoute := r.Group("/auth")
	authEmployeeRoutes(&authRoute, db)
}

func authEmployeeRoutes(route *fiber.Router, db *gorm.DB) {
	r := *route

	empRepo := repositories.NewEmployeeRepository(db)
	empApp := application.NewEmployeeApp(empRepo)

	empRoleRepo := repositories.NewEmployeeRoleRepository(db)
	empRoleApp := application.NewEmployeeRoleApp(empRoleRepo)

	authEmpApp := application.NewAuthEmployeeApp(empApp, empRoleApp)
	authEmpHandler := handlers.NewAuthEmployeeHandler(authEmpApp)

	authEmpRoute := r.Group("/employee")
	{
		authEmpRoute.Post("/register", authEmpHandler.RegisterAdmin)
		authEmpRoute.Post("/login", authEmpHandler.LoginEmployee)
	}
}

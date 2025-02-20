package routes

import (
	"thirawoot/in2forest_shop_backend/internal/adapters/database/repositories"
	"thirawoot/in2forest_shop_backend/internal/adapters/http/handlers"
	"thirawoot/in2forest_shop_backend/internal/adapters/http/middlewares"
	"thirawoot/in2forest_shop_backend/internal/application"
	"thirawoot/in2forest_shop_backend/internal/config"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(route *fiber.Router, db *gorm.DB) {
	env := config.LoadEnv()
	r := *route

	jwt := middlewares.NewAuthMiddleware(env.Secret)

	userRoute := r.Group("/user", jwt)
	userEmployeeRoute(&userRoute, db)
}

func userEmployeeRoute(route *fiber.Router, db *gorm.DB) {
	r := *route

	repo := repositories.NewEmployeeRepository(db)
	app := application.NewEmployeeApp(repo)
	handler := handlers.NewEmployeeHandler(app)

	userEmpRoute := r.Group("/employee", middlewares.RoleMiddleware(constants.Role.Admin))
	{
		userEmpRoute.Get("/profile", handler.GetProfile)
	}
}

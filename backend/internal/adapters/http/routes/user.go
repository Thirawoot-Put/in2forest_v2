package routes

import (
	"thirawoot/in2forest_shop_backend/internal/adapters/database/repositories"
	"thirawoot/in2forest_shop_backend/internal/adapters/http/handlers"
	"thirawoot/in2forest_shop_backend/internal/adapters/http/middlewares"
	"thirawoot/in2forest_shop_backend/internal/application"
	"thirawoot/in2forest_shop_backend/internal/config"

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
	hander := handlers.NewEmployeeHandler(app)

	userEmpRoute := r.Group("/employee")
	{
		userEmpRoute.Get("/profile", hander.GetProfile)
	}
}

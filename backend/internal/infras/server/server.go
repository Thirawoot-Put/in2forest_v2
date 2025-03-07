package server

import (
	"fmt"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"thirawoot/in2forest_shop_backend/internal/adapters/http/routes"
	"thirawoot/in2forest_shop_backend/internal/config"
	"thirawoot/in2forest_shop_backend/internal/infras/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
}

func AppServer() *Server {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	return &Server{app: app}
}

func (s *Server) routes(prefix string, version string, db *gorm.DB) {
	url := fmt.Sprintf("/%s/%s", prefix, version)

	mainGroup := s.app.Group(url)

	mainGroup.Use(healthcheck.New())

	// place other route
	routes.EmployeeRoleRoutes(&mainGroup, db)
	routes.AuthRoutes(&mainGroup, db)
	routes.UserRoutes(&mainGroup, db)
}

func (s *Server) Start() {
	s.app.Use(fiberlog.New())

	env := config.LoadEnv()

	db := database.Connect(env)

	addr := fmt.Sprintf(":%s", env.Port)

	s.routes(env.Prefix, env.Version, db)

	err := s.app.Listen(addr)
	if err != nil {
		log.Errorf("Failed to listen at %s: %v", addr, err)
	}
}

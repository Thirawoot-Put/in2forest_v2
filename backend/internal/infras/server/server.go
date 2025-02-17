package server

import (
	"fmt"
	"thirawoot/in2forest_shop_backend/internal/adapters/handlers"
	"thirawoot/in2forest_shop_backend/internal/config"
	"thirawoot/in2forest_shop_backend/internal/infras/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

func (s *Server) routes(prefix string, version string) {
	url := fmt.Sprintf("/%s/%s", prefix, version)
	mainGroup := s.app.Group(url)

	mainGroup.Get("/health", handlers.HealthCheck)

	// place other route
}

func (s *Server) Start() {
	env := config.LoadEnv()

	// connect db
	database.Connect(env)

	addr := fmt.Sprintf(":%s", env.Port)

	s.routes(env.Prefix, env.Version)

	err := s.app.Listen(addr)
	if err != nil {
		log.Errorf("Failed to listen at %s: %v", addr, err)
	}
}

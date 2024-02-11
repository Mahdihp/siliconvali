package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"siliconvali/config"
	"siliconvali/delivery/httpserver/userhandler"
	"siliconvali/services/authservice"
	"siliconvali/services/userservice"
)

type Server struct {
	config      config.AppConfiguration
	userHandler userhandler.Handler
	Router      *fiber.App
}

func New(config config.AppConfiguration, authSvc authservice.AuthService, userSvc userservice.UserService) Server {

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Silicon Ali",
		AppName:       "Test App v1.0.1",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	return Server{
		Router:      app,
		config:      config,
		userHandler: userhandler.New(config.AuthConfig, authSvc, userSvc),
	}
}
func (s Server) Serve() {

	s.Router.Get("/health-check", s.healthCheck)
	s.userHandler.SetRoutes(s.Router)

	if err := s.Router.Listen(s.config.Server.Host + s.config.Server.Port); err != nil {
		fmt.Println("router start error", err)
	}
}

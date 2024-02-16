package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"siliconvali/config"
	"siliconvali/delivery/httpserver/userhandler"
	"siliconvali/services/authservice"
	"siliconvali/services/userservice"
	"siliconvali/validator/uservalidator"
	"time"
)

type Server struct {
	config      config.AppConfiguration
	userHandler userhandler.Handler
	Router      *fiber.App
}

func New(config config.AppConfiguration, authSvc authservice.AuthService, userSvc userservice.UserService, userValidator uservalidator.Validator) Server {

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

	app.Use(requestid.New(requestid.Config{
		Header: "X-Custom-Header",
		Generator: func() string {
			return "static-id"
		},
	}))

	app.Use(recover.New())
	app.Use(healthcheck.New())

	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			err := c.Status(200).JSON(fiber.Map{"time": time.Now().String()})
			if err != nil {
				return false
			}
			return true
		},
		LivenessEndpoint: "/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			err := c.Status(200).JSON(fiber.Map{"time": time.Now().String()})
			if err != nil {
				return false
			}
			return true
		},
		ReadinessEndpoint: "/ready",
	}))

	return Server{
		Router:      app,
		config:      config,
		userHandler: userhandler.New(config.AuthConfig, authSvc, userSvc, userValidator),
	}
}
func (s Server) Serve() {

	//s.Router.Get("/health-check", s.healthCheck)
	s.userHandler.SetRoutes(s.Router)

	if err := s.Router.Listen(s.config.Server.Host + s.config.Server.Port); err != nil {
		fmt.Println("router start error", err)
	}
}

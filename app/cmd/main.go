package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	_ "github.com/lib/pq"
	"siliconvali/config"
)

var AppConfig config.AppConfiguration

func init() {
	AppConfig = config.Initialize()
	//fmt.Println(AppConfig)
}
func main() {

	//dbClient := postgres.InitAndDataSeeder(AppConfig)
	//ctx := context.Background()
	//count, _ := dbClient.Role.Query().Count(ctx)
	//fmt.Println("Role.Query(): ", count)

	//StartServer()
}

func StartServer() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Silicon Ali",
		AppName:       "Test App v1.0.1",
	})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))

	app.Get("/Add", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Listen(AppConfig.Server.Host + AppConfig.Server.Port)
}

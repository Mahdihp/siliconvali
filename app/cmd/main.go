package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	_ "github.com/lib/pq"
	"siliconvali/config"
	"siliconvali/dto"
	"siliconvali/repository/postgres"
	"siliconvali/repository/user_repository"
)

var AppConfig config.AppConfiguration

func init() {
	AppConfig = config.Initialize()

	//fmt.Println(AppConfig)
}
func main() {

	dbClient := postgres.New(AppConfig.DbConfig)
	//ctx := context.Background()
	//count, _ := dbClient.Role.Query().Count(ctx)
	//fmt.Println("Role.Query(): ", count)
	repositoryImpl := user_repository.New(dbClient)
	user, _ := repositoryImpl.GetAll(context.Background(), dto.GetAllUserRequest{PageIndex: 1, PageSize: 2})
	fmt.Printf("service --------: %+v \n", user)

	//service := userservice.New(repositoryImpl)
	//
	//username, _ := service.GetByMobile("mahdi", "admin")
	//fmt.Println("service --------: ", username)

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

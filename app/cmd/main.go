package main

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
	_ "github.com/lib/pq"
	"siliconvali/config"
	"siliconvali/ent"
	"siliconvali/ent/role"
)

var AppConfig config.AppConfiguration

func init() {
	AppConfig = config.Initialize()
	//fmt.Println(AppConfig)
}
func main() {

	//ConnetionString := "postgres://postgres:postgres@localhost:5432/postgres"
	ConnetionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		AppConfig.DbConfig.Host, AppConfig.DbConfig.DbPort, AppConfig.DbConfig.Username, AppConfig.DbConfig.DbName, AppConfig.DbConfig.Password)

	//fmt.Println("ConnetionString:", ConnetionString)

	client, err := ent.Open(dialect.Postgres, ConnetionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	ctx := context.Background()
	//if err := client.Schema.Create(ctx); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	//newRole, errSql := client.Role.Create().
	//	SetName("ADMIN").
	//	Save(ctx)
	//
	//fmt.Println("Role Error", errSql)
	//fmt.Println("Role Result", newRole)

	roleFund, err := client.Role.Query().Where(sql.FieldEQ(role.FieldID, 2)).First(ctx)

	a8m, errSql := client.User.Create().
		SetUsername("mahdi3").
		SetFirstname("Mahdi").
		SetLastname("Hosseinpour").
		SetMobile("09339466051").
		//AddRoleIDs(newRole.ID).
		AddRoles(roleFund).
		Save(ctx)

	fmt.Println("Role Found", roleFund)
	fmt.Println("User Error", errSql)
	fmt.Println("User Result", a8m)

	defer client.Close()
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

package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
	"siliconvali/config"
	"siliconvali/delivery/httpserver"
	"siliconvali/repositories/postgres"
	"siliconvali/repositories/user_repository"
	"siliconvali/services/authservice"
	"siliconvali/services/userservice"
	"siliconvali/validator/uservalidator"
)

var AppConfig config.AppConfiguration

func init() {
	AppConfig = config.Initialize()

	//fmt.Println(AppConfig)
}
func main() {

	authService, userService, uv := setupServices(AppConfig)
	server := httpserver.New(AppConfig, authService, userService, uv)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, AppConfig.AuthConfig.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.ShutdownWithTimeout(AppConfig.AuthConfig.GracefulShutdownTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	fmt.Println("received interrupt signal, shutting down gracefully..")
	<-ctxWithTimeout.Done()
}

func setupServices(cfg config.AppConfiguration) (
	authservice.AuthService,
	userservice.UserService, uservalidator.Validator) {

	postgresqlDB := postgres.New(cfg.DbConfig)

	authSvc := authservice.New(cfg.AuthConfig)

	userRepositoryImpl := user_repository.New(postgresqlDB)
	userService := userservice.New(authSvc, userRepositoryImpl)
	uV := uservalidator.New(userRepositoryImpl)

	return authSvc, userService, uV

}

package userhandler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"siliconvali/config"
	"siliconvali/dto"
	"siliconvali/services/authservice"
	"siliconvali/services/userservice"
)

type Handler struct {
	authConfig config.AuthConfig
	authSvc    authservice.AuthService
	userSvc    userservice.UserService
}

func New(authConfig config.AuthConfig, authSvc authservice.AuthService,
	userSvc userservice.UserService) Handler {
	return Handler{
		authConfig: authConfig,
		authSvc:    authSvc,
		userSvc:    userSvc,
	}
}
func (h Handler) userLogin(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}
	//TODO: validate request

	resp, err := h.userSvc.Login(context.Background(), req)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(resp)

}

func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/users")

	userGroup.Post("/login", h.userLogin)
	//userGroup.POST("/register", h.userRegister)
}

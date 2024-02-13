package userhandler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"siliconvali/config"
	"siliconvali/dto"
	"siliconvali/pkg/httpmsg"
	"siliconvali/services/authservice"
	"siliconvali/services/userservice"
	"siliconvali/validator/uservalidator"
)

type Handler struct {
	authConfig    config.AuthConfig
	authSvc       authservice.AuthService
	userSvc       userservice.UserService
	userValidator uservalidator.Validator
}

func New(authConfig config.AuthConfig, authSvc authservice.AuthService,
	userSvc userservice.UserService, userValidator uservalidator.Validator) Handler {
	return Handler{
		authConfig:    authConfig,
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
func (h Handler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/users")

	userGroup.Post("/login", h.userLogin)
	userGroup.Post("/all", h.getAll)
	userGroup.Post("/register", h.userRegister)
}
func (h Handler) userRegister(c *fiber.Ctx) error {
	var req dto.UserInsertRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	if fieldErrors, err := h.userValidator.ValidateRegisterRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		dataValidatore := dto.NewBaseResponse(code).
			WithMessage(msg).
			WithValidate(fieldErrors)
		return c.Status(code).JSON(dataValidatore)
	}

	resp, err := h.userSvc.Register(c.UserContext(), req)
	if err != nil {
		response := dto.NewBaseResponse(http.StatusBadRequest).
			WithError([]string{err.Error()}).
			WithMessage(err.Error()).
			WithOperation(c.Path())
		return fiber.NewError(http.StatusBadRequest, response.ToJson())
	}

	data := dto.NewBaseResponse(http.StatusOK).WithData(resp)
	return c.Status(http.StatusCreated).JSON(data)
}

func (h Handler) getAll(c *fiber.Ctx) error {
	var req dto.GetAllUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	resp, err := h.userSvc.GetAll(context.Background(), req)
	if err != nil {
		response := dto.NewBaseResponse(http.StatusBadRequest).
			WithError([]string{err.Error()}).
			WithMessage(err.Error()).
			WithOperation(c.Path())
		return fiber.NewError(http.StatusBadRequest, response.ToJson())
	}

	data := dto.NewBaseResponse(http.StatusOK).
		WithData(resp)

	return c.Status(http.StatusOK).JSON(data)

}
func (h Handler) userLogin(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}
	//TODO: validate request

	resp, err := h.userSvc.Login(c.UserContext(), req)
	if err != nil {
		response := dto.NewBaseResponse(http.StatusBadRequest).
			WithError([]string{err.Error()}).
			WithMessage(err.Error()).
			WithOperation(c.Path())
		return fiber.NewError(http.StatusBadRequest, response.ToJson())
	}
	return c.Status(http.StatusOK).JSON(resp)
}

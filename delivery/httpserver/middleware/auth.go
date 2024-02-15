package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"siliconvali/config"
	"siliconvali/services/authservice"
)

type AuthMiddleware struct {
	ctx     *fiber.Ctx
	Service authservice.AuthService
	Config  config.AuthConfig
}

func NewAuthMiddleware(ctx *fiber.Ctx) AuthMiddleware {
	return AuthMiddleware{
		ctx: ctx,
	}
}
func (Auth AuthMiddleware) AddService(service authservice.AuthService, config config.AuthConfig) AuthMiddleware {
	Auth.Config = config
	Auth.Service = service
	return Auth
}
func (Auth AuthMiddleware) Authorization() error {
	key := Auth.ctx.Get("Authorization")
	claims, err := Auth.Service.ParseToken(key)

	fmt.Println(err)
	fmt.Println(claims)

	return err
}

//func customKeyFunc(service authservice.AuthService, key string) fiber.Handler {
//	return func(t *jwt.Token) (interface{}, error) {
//		claims, err := service.ParseToken(key)
//		if err != nil {
//			return nil, err
//		}
//
//		return claims, nil
//	}
//}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

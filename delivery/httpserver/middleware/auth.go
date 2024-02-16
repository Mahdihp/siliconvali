package middleware

import (
	"fmt"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"siliconvali/config"
	"siliconvali/services/authservice"
)

func Protected(config config.AuthConfig, svc authservice.AuthService) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.SignKey)},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			jwtToken := ctx.Get("Authorization")

			_, err := svc.ParseToken(jwtToken)
			if err != nil {
				return err
			}
			//fmt.Println(claims)
			//claims := token.Claims.(*authservice.Claims)

			//fmt.Println(claims)

			return ctx.Next()
		},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func jwtSuccess(c *fiber.Ctx) error {
	jwtToken := c.Get("Authorization")
	fmt.Println(jwtToken)

	token, err := jwt.ParseWithClaims(jwtToken, &authservice.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("jwt_secret"), nil
	})

	if err != nil {
		return err
	}

	claims := token.Claims.(*authservice.Claims)

	fmt.Println(claims)
	//var user model.User
	//database.DB.Where("username = ?", claims.Issuer).First(&user)

	//c.SetUserContext(context.WithValue(c.UserContext(), "user", claims))

	return c.Next()
}

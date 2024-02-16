package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"siliconvali/pkg/claim"
	"siliconvali/services/authservice"
	"siliconvali/services/userservice"
)

func AccessCheck(userService userservice.UserService, svc authservice.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := claim.GetClaimsFromEchoContext(c, svc)

		//fmt.Println(claims)
		userInfo, err := userService.GetById(c.UserContext(), claims.UserId)
		fmt.Println(err)
		fmt.Println(userInfo)

		//return c.Status(500).JSON(fiber.Map{
		//	"message": errmsg.ErrorMsgSomethingWentWrong,
		//})
		return c.Next()
	}
}

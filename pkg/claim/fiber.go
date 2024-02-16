package claim

import (
	"github.com/gofiber/fiber/v2"
	"siliconvali/services/authservice"
)

func GetClaimsFromEchoContext(ctx *fiber.Ctx, svc authservice.AuthService) *authservice.Claims {
	jwtToken := ctx.Get("Authorization")
	claims, err := svc.ParseToken(jwtToken)
	if err != nil {
		return nil
	}
	return claims
}

package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"siliconvali/dto"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId int64          `json:"user_id"`
	Roles  []dto.RoleInfo `json:"roles"`
}

func (c Claims) Valid() error {
	return c.Valid()
}

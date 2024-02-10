package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"siliconvali/config"
	"siliconvali/dto"
	"strings"
	"time"
)

type AuthService struct {
	config config.AuthConfig
}

func New(cfg config.AuthConfig) AuthService {
	return AuthService{
		config: cfg,
	}
}
func (s AuthService) CreateAccessToken(user dto.UserInfo) (string, error) {
	return s.createToken(user.UserId, user.Roles, s.config.AccessSubject, s.config.AccessExpirationTime)
}

func (s AuthService) CreateRefreshToken(user dto.UserInfo) (string, error) {
	return s.createToken(user.UserId, user.Roles, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s AuthService) ParseToken(bearerToken string) (*Claims, error) {
	//https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType

	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
func (s AuthService) createToken(userID int64, roles []dto.RoleInfo, subject string, expireDuration time.Duration) (string, error) {
	// create a signer for rsa 256
	// TODO - replace with rsa 256 RS256 - https://github.com/golang-jwt/jwt/blob/main/http_example_test.go

	// set our claims
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserId: userID,
		Roles:  roles,
	}

	// TODO - add sign method to config
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

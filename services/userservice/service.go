package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"siliconvali/dto"
	"siliconvali/ent"
)

type Repository interface {
	Register(reqDto dto.RegisterRequest) (ent.User, error)
	GetUsername(username string) (ent.User, error)
	GetUserByID(ctx context.Context, userID uint) (ent.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user ent.User) (string, error)
	CreateRefreshToken(user ent.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{auth: authGenerator, repo: repo}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

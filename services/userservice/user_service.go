package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"siliconvali/dto"
	"siliconvali/ent"
)

type UserServiceRepository interface {
	Register(u dto.UserInsertRequest) (dto.UserInsertResponse, error)
	GetUserByPhoneNumber(phoneNumber string) (dto.UserInfo, error)
	GetUserByID(ctx context.Context, userId uint) (dto.UserInfo, error)
}

type AuthGenerator interface {
	CreateAccessToken(user ent.User) (string, error)
	CreateRefreshToken(user ent.User) (string, error)
}

type UserServiceRepositoryImpl struct {
	auth AuthGenerator
	repo UserServiceRepository
}

func New(authGenerator AuthGenerator, repo UserServiceRepository) UserServiceRepositoryImpl {
	return UserServiceRepositoryImpl{
		auth: authGenerator,
		repo: repo,
	}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

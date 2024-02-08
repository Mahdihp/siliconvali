package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"siliconvali/dto"
	"siliconvali/ent"
)

type UserServiceRepositort interface {
	Register(u dto.UserInsertRequest) (dto.UserInsertResponse, error)
	GetUserByPhoneNumber(phoneNumber string) (dto.UserInfo, error)
	GetUserByID(ctx context.Context, userId uint) (dto.UserInfo, error)
}

type AuthGenerator interface {
	CreateAccessToken(user ent.User) (string, error)
	CreateRefreshToken(user ent.User) (string, error)
}

type UserServiceRepositortImpl struct {
	auth AuthGenerator
	repo UserServiceRepositort
}

func New(authGenerator AuthGenerator, repo UserServiceRepositort) UserServiceRepositortImpl {
	return UserServiceRepositortImpl{
		auth: authGenerator,
		repo: repo,
	}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

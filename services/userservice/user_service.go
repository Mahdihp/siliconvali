package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"siliconvali/dto"
)

type UserService interface {
	Register(u dto.UserInsertRequest) (dto.UserInsertResponse, error)
	GetUserByPhoneNumber(phoneNumber string) (dto.UserInfo, error)
	GetUserByID(ctx context.Context, userId uint) (dto.UserInfo, error)
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

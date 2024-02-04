package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"siliconvali/ent"
	"siliconvali/repository/postgresuser"
)

type UserService struct {
	repo postgresuser.UserRepository
}

func New(repo postgresuser.UserRepository) UserService {
	return UserService{repo: repo}
}

func (us UserService) GetByUsername(username string, password string) (ent.User, error) {
	getUsername, err := us.repo.GetByUsername(context.Background(), username)
	return getUsername, err
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

package userservice

import (
	"fmt"
	"siliconvali/dto"
	"siliconvali/ent"
	richerror "siliconvali/pkg"
)

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	// TODO - we should verify phone number by verification code

	// TODO - replace md5 with bcrypt
	user := ent.User{
		ID:       0,
		Username: req.Username,
		Password: getMD5Hash(req.Password),
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	return dto.RegisterResponse{User: dto.UserInfo{
		ID:        createdUser.ID,
		Username:  createdUser.Username,
		FirstName: *createdUser.Firstname,
		LastName:  *createdUser.Lastname,
	}}, nil
}
func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	const op = "userservice.Login"

	user, err := s.repo.GetUsername(req.Username)
	if err != nil {
		return dto.LoginResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"username": req.Username})
	}

	if user.Username != getMD5Hash(req.Password) {
		return dto.LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return dto.LoginResponse{
		User: dto.UserInfo{
			ID:           user.ID,
			Username:     user.Username,
			FirstName:    *user.Firstname,
			LastName:     *user.Lastname,
			Mobile:       user.Mobile,
			NationalCode: user.NationalCode,
			Active:       user.Active,
			Address:      *user.Address,
		},
		Tokens: dto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken},
	}, nil
}

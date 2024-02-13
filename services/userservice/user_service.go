package userservice

import (
	"context"
	"fmt"
	"siliconvali/dto"
	richerror "siliconvali/pkg"
	"siliconvali/pkg/errmsg"
	"siliconvali/pkg/util"
	"siliconvali/repositories/user_repository"
)

type Repository interface {
	Register(ctx context.Context, u dto.UserInsertRequest) (dto.UserInfo, error)
	GetById(ctx context.Context, userId int64) (dto.UserInfo, error)
	GetAll(ctx context.Context, req dto.GetAllUserRequest) ([]dto.UserInfo, error)
	Update(ctx context.Context, req dto.UserUpdateRequest) error
}

type AuthGenerator interface {
	CreateAccessToken(user dto.UserInfo) (string, error)
	CreateRefreshToken(user dto.UserInfo) (string, error)
}

type UserService struct {
	auth AuthGenerator
	repo user_repository.UserRepository
}

func New(authGenerator AuthGenerator, repo user_repository.UserRepository) UserService {
	return UserService{
		auth: authGenerator,
		repo: repo,
	}
}

func (receiver UserService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	const op = "UserService.Login"

	user, err := receiver.repo.GetByMobile(ctx, req.Mobile)
	if err != nil {
		return dto.LoginResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"mobile": req.Mobile})
	}
	hash := util.StringToMD5Hash(req.Password)
	//fmt.Println(user.Password != hash)

	if user.Password != hash {
		return dto.LoginResponse{}, fmt.Errorf(errmsg.ErrorMsg_User_Incorrect_User_Pass)
	}

	accessToken, err := receiver.auth.CreateAccessToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	refreshToken, err := receiver.auth.CreateRefreshToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	user.Password = ""
	return dto.LoginResponse{
		User: user,
		Tokens: dto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken},
	}, nil

}
func (receiver UserService) Update(ctx context.Context, req dto.UserUpdateRequest) error {
	const op = "UserService.Update"
	err := receiver.repo.Update(ctx, req)
	if err != nil {
		return richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"req": req})
	}
	return nil
}
func (receiver UserService) GetAll(ctx context.Context, req dto.GetAllUserRequest) ([]dto.UserInfo, error) {
	const op = "UserService.GetAll"
	all, err := receiver.repo.GetAll(ctx, req)
	if err != nil {
		return []dto.UserInfo{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"req": req})
	}
	return all, nil
}
func (receiver UserService) GetById(ctx context.Context, userId int64) (dto.UserInfo, error) {
	const op = "UserService.GetUserByID"

	userInfo, err := receiver.repo.GetById(ctx, userId)
	if err != nil {
		return dto.UserInfo{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"req": userId})
	}
	return userInfo, nil
}
func (receiver UserService) Register(ctx context.Context, req dto.UserInsertRequest) (dto.UserInfo, error) {

	insert, err := receiver.repo.Insert(ctx, req)
	if err != nil {
		return dto.UserInfo{}, fmt.Errorf("unexpected error: %w", err)
	}

	return insert, nil
}

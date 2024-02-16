package user_repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"siliconvali/dto"
	"siliconvali/ent"
	"siliconvali/ent/user"
	richerror "siliconvali/pkg"
	"siliconvali/pkg/errmsg"
	"siliconvali/repositories/postgres"
)

// https://github.com/gocasts-bootcamp/gameapp/blob/main/repository/mysql/mysqluser/user.go

type UserRepository interface {
	Insert(ctx context.Context, req dto.UserInsertRequest) (dto.UserInfo, error)
	Update(ctx context.Context, req dto.UserUpdateRequest) error
	DeleteById(ctx context.Context, userId int64) error
	GetByMobile(ctx context.Context, mobile string) (dto.UserInfo, error)
	IsMobileUnique(ctx context.Context, mobile string) (bool, error)
	GetById(ctx context.Context, userId int64) (dto.UserInfo, error)
	GetAll(ctx context.Context, req dto.GetAllUserRequest) ([]dto.UserInfo, error)
}

type UserRepositoryImpl struct {
	conn *postgres.PostgresqlDB
}

func New(conn *postgres.PostgresqlDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		conn: conn,
	}
}

func (userRepo *UserRepositoryImpl) IsMobileUnique(ctx context.Context, mobile string) (bool, error) {
	const op = "user_repository.IsMobileUnique"

	found, err := userRepo.conn.Conn().User.Query().
		Where(sql.FieldEQ(user.FieldMobile, mobile)).
		Exist(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return true, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}
		return false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return found, nil
}

func (userRepo *UserRepositoryImpl) DeleteById(ctx context.Context, userId int64) error {
	const op = "user_repository.DeleteById"

	userFound, err := userRepo.conn.Conn().User.Query().
		Where(user.And(sql.FieldEQ(user.FieldID, userId), sql.FieldEQ(user.FieldDeleted, false))).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}
		//var not *ent.NotFoundError
		//if errors.As(err, &not) {
		//
		//}

		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	_, errDelete := userFound.Update().SetDeleted(true).Save(ctx)
	return errDelete
}

func (userRepo *UserRepositoryImpl) GetById(ctx context.Context, userId int64) (dto.UserInfo, error) {
	const op = "user_repository.GetById"

	userFound, err := userRepo.conn.Conn().User.Query().
		Where(sql.FieldEQ(user.FieldID, userId)).
		WithRoles().
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return dto.UserInfo{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return dto.UserInfo{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return UserToUserInfo(userFound), nil
}

func (userRepo *UserRepositoryImpl) GetAll(ctx context.Context, req dto.GetAllUserRequest) ([]dto.UserInfo, error) {
	const op = "user_repository.GetAll"

	users, err := userRepo.conn.Conn().User.Query().
		Order(ent.Desc(user.FieldCreatedAt)).
		Limit(req.PageSize).
		Offset(req.PageIndex).
		WithRoles().
		All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return []dto.UserInfo{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return []dto.UserInfo{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return UsersToUserInfos(users), nil
}

func (userRepo *UserRepositoryImpl) Insert2(ctx context.Context, u ent.User) (ent.User, error) {

	newUser, err := userRepo.conn.Conn().User.Create().
		SetMobile(u.Mobile).
		SetFirstname(*u.Firstname).
		SetLastname(*u.Lastname).
		SetNationalCode(u.NationalCode).
		SetAddress(*u.Address).
		SetActive(true).
		AddRoles(u.Edges.Roles...).
		Save(ctx)

	if err != nil {
		return ent.User{}, fmt.Errorf("can't execute command: %w", err)
	}

	return *newUser, nil
}
func (userRepo *UserRepositoryImpl) Insert(ctx context.Context, u dto.UserInsertRequest) (dto.UserInfo, error) {

	newUser, err := userRepo.conn.Conn().User.Create().
		SetMobile(u.Mobile).
		SetFirstname(u.FirstName).
		SetLastname(u.LastName).
		SetNationalCode(u.NationalCode).
		SetAddress(u.Address).
		SetActive(false).
		AddRoleIDs(u.RoleId).
		Save(ctx)

	if err != nil {
		return dto.UserInfo{}, fmt.Errorf("can't execute command: %w", err)
	}

	return UserToUserInfo(newUser), nil
}

func (userRepo *UserRepositoryImpl) GetByMobile(ctx context.Context, mobile string) (dto.UserInfo, error) {
	const op = "user_repository.GetByMobile"

	userFound, err := userRepo.conn.Conn().User.Query().
		Where(sql.FieldEQ(user.FieldMobile, mobile)).
		WithRoles().
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return dto.UserInfo{}, richerror.New(op).
				WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}

		return dto.UserInfo{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return UserToUserInfo(userFound), nil
}

func (userRepo UserRepositoryImpl) Update(ctx context.Context, req dto.UserUpdateRequest) error {
	const op = "user_repository.Update"

	userFound, err := userRepo.conn.Conn().User.Query().
		Where(sql.FieldEQ(user.FieldID, req.UserId)).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	_, errUpdate := userFound.Update().SetMobile(req.Mobile).
		SetFirstname(req.FirstName).
		SetLastname(req.LastName).
		SetNationalCode(req.NationalCode).
		SetAddress(req.Address).
		SetMobile(req.Mobile).
		AddRoleIDs(req.RoleId).
		Save(ctx)

	return errUpdate
}

func UsersToUserInfos(users []*ent.User) []dto.UserInfo {
	var userInfos []dto.UserInfo
	for _, u := range users {

		userInfos = append(userInfos, UserToUserInfo(u))
	}
	return userInfos
}
func UserToUserInfo(user *ent.User) dto.UserInfo {
	userInfo := dto.UserInfo{
		UserId:       user.ID,
		Mobile:       user.Mobile,
		Password:     user.Password,
		NationalCode: user.NationalCode,
		Active:       user.Active,
		Deleted:      user.Deleted,
	}
	for _, role := range user.Edges.Roles {
		userInfo.Roles = append(userInfo.Roles, dto.RoleInfo{
			Id:          role.ID,
			Name:        role.Name,
			Description: *role.Description,
		})
	}
	if user.Address != nil {
		userInfo.Address = *user.Address
	}
	if user.Firstname != nil {
		userInfo.FirstName = *user.Firstname
	}
	if user.Lastname != nil {
		userInfo.LastName = *user.Lastname
	}
	return userInfo
}

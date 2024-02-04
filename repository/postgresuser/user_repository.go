package postgresuser

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"siliconvali/dto"
	"siliconvali/ent"
	"siliconvali/ent/user"
	richerror "siliconvali/pkg"
	"siliconvali/pkg/errmsg"
	"siliconvali/repository/postgres"
)

// https://github.com/gocasts-bootcamp/gameapp/blob/main/repository/mysql/mysqluser/user.go

type UserRepositoryImpl struct {
	conn *postgres.PostgresqlDB
}

func New(conn *postgres.PostgresqlDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		conn: conn,
	}
}

type UserRepository interface {
	Insert(req dto.InsertRequest, ctx context.Context) (ent.User, error)
	Update(req dto.UpdateRequest, ctx context.Context) (dto.UserInfo, error)
	GetByUsername(ctx context.Context, username string) (ent.User, error)
	GetUserByID(ctx context.Context, userId int64) (ent.User, error)
}

func (userRepo *UserRepositoryImpl) Insert(u dto.InsertRequest, ctx context.Context) (ent.User, error) {

	newUser, err := userRepo.conn.Conn().User.Create().
		SetUsername(u.Username).
		SetFirstname(u.FirstName).
		SetLastname(u.LastName).
		SetNationalCode(u.NationalCode).
		SetMobile(u.Mobile).
		SetAddress(u.Address).SetActive(false).Save(ctx)

	if err != nil {
		return ent.User{}, fmt.Errorf("can't execute command: %w", err)
	}

	return *newUser, nil
}

func (userRepo *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (ent.User, error) {
	const op = "postgresuser.GetByUsername"

	userFound, err := userRepo.conn.Conn().User.Query().
		Where(sql.FieldEQ(user.FieldUsername, username)).
		First(ctx)

	if err != nil {
		var not *ent.NotFoundError
		if errors.As(err, &not) {
			return ent.User{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return ent.User{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return *userFound, nil
}

func (userRepo *UserRepositoryImpl) GetUserByID(ctx context.Context, userId int64) (ent.User, error) {
	const op = "postgresuser.GetUserByID"

	userFound, err := userRepo.conn.Conn().User.Query().
		Where(sql.FieldEQ(user.FieldID, userId)).
		First(ctx)

	if err != nil {
		var not *ent.NotFoundError
		if errors.As(err, &not) {
			return ent.User{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return ent.User{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return *userFound, nil
}

func (userRepo UserRepositoryImpl) Update(req dto.UpdateRequest, ctx context.Context) (dto.UserInfo, error) {
	const op = "postgresuser.Update"

	userFound, err := userRepo.conn.Conn().User.Query().
		Where(sql.FieldEQ(user.FieldID, req.UserId)).
		First(ctx)

	if err != nil {
		var not *ent.NotFoundError
		if errors.As(err, &not) {
			return dto.UserInfo{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return dto.UserInfo{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	save, _ := userFound.Update().SetUsername(req.Username).
		SetFirstname(req.FirstName).
		SetLastname(req.LastName).
		SetNationalCode(req.NationalCode).
		SetMobile(req.Mobile).
		SetAddress(req.Address).
		SetMobile(req.Mobile).
		Save(ctx)

	return UserToUserInfo(save), nil
}
func UserToUserInfo(user *ent.User) dto.UserInfo {
	return dto.UserInfo{
		Id:           user.ID,
		Username:     user.Username,
		FirstName:    *user.Firstname,
		LastName:     *user.Lastname,
		NationalCode: user.NationalCode,
		Mobile:       user.Mobile,
		Address:      *user.Address,
		Active:       user.Active,
	}
}

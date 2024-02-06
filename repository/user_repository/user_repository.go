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
	Insert(ctx context.Context, req dto.InsertRequest) (dto.UserInfo, error)
	Update(ctx context.Context, req dto.UpdateRequest) error
	DeleteById(ctx context.Context, userId int64) error
	GetByUsername(ctx context.Context, username string) (dto.UserInfo, error)
	GetById(ctx context.Context, userId int64) (dto.UserInfo, error)
	GetAll(ctx context.Context, req dto.GetAllRequest) ([]dto.UserInfo, error)
}

func (userRepo *UserRepositoryImpl) DeleteById(ctx context.Context, userId int64) error {
	const op = "postgresuser.DeleteById"

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
	const op = "postgresuser.GetById"

	userFound, err := userRepo.conn.Conn().User.Query().
		Where(sql.FieldEQ(user.FieldID, userId)).
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

func (userRepo *UserRepositoryImpl) GetAll(ctx context.Context, req dto.GetAllRequest) ([]dto.UserInfo, error) {
	const op = "postgresuser.GetAll"

	users, err := userRepo.conn.Conn().User.Query().
		Order(ent.Desc(user.FieldCreatedAt)).
		Limit(req.PageSize).
		Offset(req.PageIndex).
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

func (userRepo *UserRepositoryImpl) Insert(ctx context.Context, u dto.InsertRequest) (dto.UserInfo, error) {

	newUser, err := userRepo.conn.Conn().User.Create().
		SetUsername(u.Username).
		SetFirstname(u.FirstName).
		SetLastname(u.LastName).
		SetNationalCode(u.NationalCode).
		SetMobile(u.Mobile).
		SetAddress(u.Address).SetActive(false).Save(ctx)

	if err != nil {
		return dto.UserInfo{}, fmt.Errorf("can't execute command: %w", err)
	}

	return UserToUserInfo(newUser), nil
}

func (userRepo *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (dto.UserInfo, error) {
	const op = "postgresuser.GetByUsername"

	userFound, err := userRepo.conn.Conn().User.Query().
		Where(sql.FieldEQ(user.FieldUsername, username)).
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

func (userRepo UserRepositoryImpl) Update(ctx context.Context, req dto.UpdateRequest) error {
	const op = "postgresuser.Update"

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

	_, errUpdate := userFound.Update().SetUsername(req.Username).
		SetFirstname(req.FirstName).
		SetLastname(req.LastName).
		SetNationalCode(req.NationalCode).
		SetMobile(req.Mobile).
		SetAddress(req.Address).
		SetMobile(req.Mobile).
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
		Id:           user.ID,
		Username:     user.Username,
		NationalCode: user.NationalCode,
		Mobile:       user.Mobile,
		Active:       user.Active,
		Deleted:      user.Deleted,
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

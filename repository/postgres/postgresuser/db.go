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

type DB struct {
	conn *postgres.PostgresqlDB
}

func New(conn *postgres.PostgresqlDB) *DB {
	return &DB{
		conn: conn,
	}
}

// https://github.com/gocasts-bootcamp/gameapp/blob/main/repository/mysql/mysqluser/user.go

func (d *DB) IsUsernameUnique(username string) (bool, error) {
	const op = "postgres.IsUsernameUnique"
	ctx := context.Background()

	found, err := d.conn.Conn().User.Query().Where(sql.FieldEQ(user.FieldUsername, username)).Exist(ctx)
	if err != nil {
		return false, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return found, nil
}
func (d *DB) Register(u dto.RegisterRequest) (ent.User, error) {

	ctx := context.Background()
	newUser, err := d.conn.Conn().User.Create().
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
func (d *DB) GetUserByID(ctx context.Context, userId int64) (ent.User, error) {
	const op = "postgres.GetUserByID"

	userFound, err := d.conn.Conn().User.Query().
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

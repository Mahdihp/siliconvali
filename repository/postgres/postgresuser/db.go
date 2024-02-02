package postgresuser

import (
	"context"
	"entgo.io/ent/dialect/sql"
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

func (d *DB) IsUsernameUnique(username string) (bool, error) {
	const op = "postgres.IsUsernameUnique"
	ctx := context.Background()

	found, err := d.conn.Conn().User.Query().Where(sql.FieldEQ(user.FieldUsername, username)).Exist(ctx)
	if err != nil {
		return false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return found, nil
}

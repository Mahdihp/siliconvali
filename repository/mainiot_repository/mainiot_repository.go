package mainiot_repository

import (
	"context"
	"siliconvali/dto"
	"siliconvali/repository/postgres"
)

type MainIotRepositoryImpl struct {
	conn *postgres.PostgresqlDB
}

func New(conn *postgres.PostgresqlDB) *MainIotRepositoryImpl {
	return &MainIotRepositoryImpl{
		conn: conn,
	}
}

type MainIotRepository interface {
	Insert(ctx context.Context, req dto.InsertRequest) (dto.UserInfo, error)
	Update(ctx context.Context, req dto.UpdateRequest) error
	DeleteById(ctx context.Context, userId int64) error
	GetById(ctx context.Context, userId int64) (dto.UserInfo, error)
	GetAll(ctx context.Context, req dto.GetAllRequest) ([]dto.UserInfo, error)
}

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
	Insert(ctx context.Context, req dto.MainIotInsertRequest) (dto.MainIotInfo, error)
	Update(ctx context.Context, req dto.MainIotUpdateRequest) error
	DeleteById(ctx context.Context, mainIotId int64) error
	GetById(ctx context.Context, mainIotId int64) (dto.MainIotInfo, error)
	GetAll(ctx context.Context, req dto.MainIotGetAllRequest) ([]dto.MainIotInfo, error)
}

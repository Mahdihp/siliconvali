package userpaymentplan_repository

import (
	"context"
	"siliconvali/dto"
	"siliconvali/repositories/postgres"
)

type UppRepositoryImpl struct {
	conn *postgres.PostgresqlDB
}

func New(conn *postgres.PostgresqlDB) *UppRepositoryImpl {
	return &UppRepositoryImpl{
		conn: conn,
	}
}

type UppRepository interface {
	Insert(ctx context.Context, req dto.UppInsertRequest) (dto.UppInfo, error)
	Update(ctx context.Context, req dto.UppUpdateRequest) error
	DeleteById(ctx context.Context, uppId int64) error
	GetById(ctx context.Context, uppId int64) (dto.UppInfo, error)
	GetAll(ctx context.Context, req dto.GetAllUppRequest) ([]dto.UppInfo, error)
}

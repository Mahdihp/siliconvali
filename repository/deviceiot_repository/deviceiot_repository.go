package deviceiot_repository

import (
	"context"
	"siliconvali/dto"
	"siliconvali/repository/postgres"
)

type DeviceIotRepositoryImpl struct {
	conn *postgres.PostgresqlDB
}

func New(conn *postgres.PostgresqlDB) *DeviceIotRepositoryImpl {
	return &DeviceIotRepositoryImpl{
		conn: conn,
	}
}

type DeviceIotRepository interface {
	Insert(ctx context.Context, req dto.DeviceIotInsertRequest) (dto.DeviceIotInfo, error)
	Update(ctx context.Context, req dto.DeviceIotUpdateRequest) error
	DeleteById(ctx context.Context, deviceiotId int64) error
	GetById(ctx context.Context, deviceiotId int64) (dto.DeviceIotInfo, error)
	GetAll(ctx context.Context, req dto.GetAllDeviceIotRequest) ([]dto.DeviceIotInfo, error)
}

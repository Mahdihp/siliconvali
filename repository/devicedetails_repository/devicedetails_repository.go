package devicedetails_repository

import (
	"context"
	"siliconvali/dto"
	"siliconvali/repository/postgres"
)

type DeviceDetailsRepositoryImpl struct {
	conn *postgres.PostgresqlDB
}

func New(conn *postgres.PostgresqlDB) *DeviceDetailsRepositoryImpl {
	return &DeviceDetailsRepositoryImpl{
		conn: conn,
	}
}

type DeviceDetailsRepository interface {
	Insert(ctx context.Context, req dto.DeviceDetailsInsertRequest) (dto.DeviceDetailsInfo, error)
	Update(ctx context.Context, req dto.DeviceDetailsUpdateRequest) error
	DeleteById(ctx context.Context, devicedetailsId int64) error
	GetById(ctx context.Context, devicedetailsId int64) (dto.DeviceDetailsInfo, error)
	GetAll(ctx context.Context, req dto.GetAllDeviceDetailsRequest) ([]dto.DeviceDetailsInfo, error)
}

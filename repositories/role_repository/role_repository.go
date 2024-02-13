package role_repository

import (
	"context"
	"siliconvali/dto"
	"siliconvali/ent"
	"siliconvali/ent/role"
	richerror "siliconvali/pkg"
	"siliconvali/pkg/errmsg"
	"siliconvali/repositories/postgres"
)

type RoleRepositoryImpl struct {
	conn *postgres.PostgresqlDB
}

func New(conn *postgres.PostgresqlDB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{
		conn: conn,
	}
}

type RoleRepository interface {
	//Insert(ctx context.Context, req dto.DeviceDetailsInsertRequest) (dto.DeviceDetailsInfo, error)
	//Update(ctx context.Context, req dto.DeviceDetailsUpdateRequest) error
	//DeleteById(ctx context.Context, devicedetailsId int64) error
	//GetById(ctx context.Context, devicedetailsId int64) (dto.DeviceDetailsInfo, error)
	GetAll(ctx context.Context, req dto.GetAllRoleRequest) ([]dto.RoleInfo, error)
}

func (receiver RoleRepositoryImpl) GetAll(ctx context.Context, req dto.GetAllRoleRequest) ([]dto.RoleInfo, error) {

	const op = "role_repository.GetAll"

	roles, err := receiver.conn.Conn().Role.Query().
		Order(ent.Asc(role.FieldID)).
		Limit(req.PageSize).
		Offset(req.PageIndex).
		All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return []dto.RoleInfo{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return []dto.RoleInfo{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	var roleInfos []dto.RoleInfo
	for _, r := range roles {
		roleInfos = append(roleInfos, dto.RoleInfo{
			Id:          r.ID,
			Name:        r.Name,
			Description: *r.Description,
		})
	}
	return roleInfos, nil

}

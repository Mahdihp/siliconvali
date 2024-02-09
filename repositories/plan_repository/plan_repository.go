package plan_repository

import (
	"context"
	"siliconvali/dto"
	"siliconvali/ent"
	"siliconvali/ent/plan"
	richerror "siliconvali/pkg"
	"siliconvali/pkg/errmsg"
	"siliconvali/repositories/postgres"
)

type PlanRepositoryImpl struct {
	conn *postgres.PostgresqlDB
}

func New(conn *postgres.PostgresqlDB) *PlanRepositoryImpl {
	return &PlanRepositoryImpl{
		conn: conn,
	}
}

type PlanRepository interface {
	//Insert(ctx context.Context, req dto.DeviceDetailsInsertRequest) (dto.DeviceDetailsInfo, error)
	//Update(ctx context.Context, req dto.DeviceDetailsUpdateRequest) error
	//DeleteById(ctx context.Context, devicedetailsId int64) error
	//GetById(ctx context.Context, devicedetailsId int64) (dto.DeviceDetailsInfo, error)
	GetAll(ctx context.Context, req dto.GetAllPlanRequest) ([]dto.PlanInfo, error)
}

func (receiver PlanRepositoryImpl) GetAll(ctx context.Context, req dto.GetAllPlanRequest) ([]dto.PlanInfo, error) {

	const op = "role_repository.GetAll"

	plans, err := receiver.conn.Conn().Plan.Query().
		Order(ent.Asc(plan.FieldID)).
		Limit(req.PageSize).
		Offset(req.PageIndex).
		All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return []dto.PlanInfo{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return []dto.PlanInfo{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	var planInfos []dto.PlanInfo
	for _, r := range plans {
		planInfos = append(planInfos, dto.PlanInfo{
			Id:          r.ID,
			Name:        r.Name,
			Price:       *r.Price,
			Period:      *r.Period,
			Active:      r.Active,
			Description: *r.Description,
		})
	}
	return planInfos, nil

}

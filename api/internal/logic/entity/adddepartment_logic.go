package entity

import (
	"context"
	"database/sql"

	"admin-server/api/internal/errorx"
	"admin-server/api/internal/model/entity"
	"admin-server/api/internal/pkg/utils"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdddepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdddepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdddepartmentLogic {
	return &AdddepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdddepartmentLogic) Adddepartment(req *types.AddDepartmentReq) error {
	userID,err:=utils.GetUserIdFromCtx(l.ctx)
	if err!=nil{
		return err
	}
	
	entityQuery:=l.svcCtx.EntityModel.SelectBuilder().Where(
		squirrel.Eq{"owner_id":userID},
	)
	count,err:=l.svcCtx.EntityModel.FindCount(l.ctx,entityQuery,"id")
	if err!=nil{
		return err
	}
	
	if count==0{
		return errorx.ErrEntityNotFound
	}
	department:=entity.Department{
		EntityId: req.EntityId,
		Status: entity.DepartmentStatusNormal,
		Name: req.Name,
		Sort: req.Sort,
	}
	if req.ParentID!=0{
		department.ParentId=sql.NullInt64{Int64: req.ParentID,Valid: true}
	}
    
	_,err=l.svcCtx.DepartmentModel.Insert(l.ctx,nil,&department)
	if err!=nil{
		return err
	}
	return nil
}

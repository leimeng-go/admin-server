package entity

import (
	"context"
	"database/sql"

	"admin-server/api/internal/errorx"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatedepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatedepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatedepartmentLogic {
	return &UpdatedepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatedepartmentLogic) Updatedepartment(req *types.UpdateDepartmentReq) error {
	department,err:=l.svcCtx.DepartmentModel.FindOne(l.ctx,req.ID)
	if err!=nil{
		return err
	}
	if department==nil{
		return errorx.ErrEntityNotFound
	}
	if department.Name!=req.Name{
		department.Name=req.Name	
	}
	if department.Sort!=req.Sort{
		department.Sort=req.Sort
	}
	if department.Status!=req.Status{
		department.Status=req.Status
	}
	if req.ParentID!=0{
		department.ParentId=sql.NullInt64{Int64: req.ParentID,Valid: true}
	}

	_,err=l.svcCtx.DepartmentModel.Update(l.ctx,nil,department)
	if err!=nil{
		return err
	}
	
	return nil
}

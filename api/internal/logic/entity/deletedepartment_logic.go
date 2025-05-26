package entity

import (
	"context"

	"admin-server/api/internal/errorx"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletedepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletedepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletedepartmentLogic {
	return &DeletedepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletedepartmentLogic) Deletedepartment(req *types.DeleteDepartmentReq) error {
    department,err:=l.svcCtx.DepartmentModel.FindOne(l.ctx,req.ID)    
	if err!=nil{
		return err
	}
	if department==nil{
		return errorx.ErrEntityNotFound
	}
	// TODO: 需要添加权限判断
	err=l.svcCtx.DepartmentModel.Delete(l.ctx,nil,req.ID)
	if err!=nil{
		return err
	}
	return nil
}

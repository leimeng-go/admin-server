package entity

import (
	"context"

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

func (l *UpdatedepartmentLogic) Updatedepartment() (resp *types.Department, err error) {
	// todo: add your logic here and delete this line

	return
}

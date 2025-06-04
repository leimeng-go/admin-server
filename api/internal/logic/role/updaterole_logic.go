package role

import (
	"context"
	"database/sql"

	"admin-server/api/internal/model/permission"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateroleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateroleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateroleLogic {
	return &UpdateroleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateroleLogic) Updaterole(req *types.UpdateRoleReq) error {
	_, err := l.svcCtx.RoleModel.Update(l.ctx, nil, &permission.Role{
		Id:          req.ID,
		Name:        req.Name,
		Description: sql.NullString{String: req.Description, Valid: true},
		Status:      req.Status,
	})
	if err != nil {
		return err
	}
	return nil
}

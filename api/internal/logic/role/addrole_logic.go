package role

import (
	"context"
	"database/sql"

	"admin-server/api/internal/model/permission"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddroleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddroleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddroleLogic {
	return &AddroleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddroleLogic) Addrole(req *types.AddRoleReq) error {
	_,err:=l.svcCtx.RoleModel.Insert(l.ctx, nil,&permission.Role{
		Name:        req.Name,
		Description: sql.NullString{String: req.Description, Valid: true},
		Status:      1,
		DeleteTime: sql.NullTime{
		},
	})
	if err != nil {
		return err
	}
	return nil 

}

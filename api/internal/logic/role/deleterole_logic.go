package role

import (
	"context"

	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteroleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteroleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteroleLogic {
	return &DeleteroleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteroleLogic) Deleterole(req *types.DeleteRoleReq) error {
	err := l.svcCtx.RoleModel.Delete(l.ctx, nil, req.ID)
	if err != nil {
		return err
	}
	return nil
}

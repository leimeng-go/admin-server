package role

import (
	"context"

	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/golang-module/carbon/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type RolelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRolelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RolelistLogic {
	return &RolelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RolelistLogic) Rolelist(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	query:=l.svcCtx.RoleModel.SelectBuilder()
    roleList, total, err := l.svcCtx.RoleModel.FindPageListByPageWithTotal(l.ctx,query,req.Page,req.PageSize,"id DESC")
	if err != nil {
		return nil, err
	}
	resp = &types.RoleListResp{
		Total: total,
		List:  make([]types.Role,0,len(roleList)),
	}
	for _, role := range roleList {
		resp.List = append(resp.List, types.Role{
			ID:          role.Id,
			Name:        role.Name,
			Description: role.Description.String,
			CreateTime:   carbon.CreateFromStdTime(role.CreateTime).ToDateTimeString(),
			UpdateTime:   carbon.CreateFromStdTime(role.UpdateTime).ToDateTimeString(),
			Status:      role.Status,
		})
	}
	
	return
}

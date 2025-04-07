package user

import (
	"context"

	"admin-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMenuLogic {
	return &GetUserMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMenuLogic) GetUserMenu() ([]*types.Menu, error) {
	// 获取当前登录用户
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, l.svcCtx.UserModel.Username)
	if err != nil {
		return nil, err
	}

	// 获取用户角色
	role, err := l.svcCtx.RoleModel.FindOneByUserId(l.ctx, user.Id)
	if err != nil {
		return nil, err
	}

	// 获取用户菜单
	menu, err := l.svcCtx.MenuModel.FindOneByRoleId(l.ctx, role.Id)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

package user

import (
	"admin-server/internal/svc"
	"admin-server/internal/types"
	"admin-server/internal/utils"
	"context"

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

func (l *GetUserMenuLogic) GetUserMenu() ([]*types.Route, error) {
	userID, err := utils.GetUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	// 获取当前登录用户
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(userID))
	if err != nil {
		return nil, err
	}
	role, err := l.svcCtx.UserRoleModel.FindOneByUserId(l.ctx, int64(user.Id))
	if err != nil {
		return nil, err
	}
	menuRoles, err := l.svcCtx.RoleMenuModel.FindMenuIDsByRoleId(l.ctx, role.RoleId)
	if err != nil {
		return nil, err
	}
	menuIDs := make([]int64, 0, len(menuRoles))
	for _, v := range menuRoles {
		menuIDs = append(menuIDs, v.MenuId)
	}
	menus, err := l.svcCtx.MenuModel.FindMenusByIDs(l.ctx, menuIDs)
	if err != nil {
		return nil, err
	}

	// 1. 构建id到Route的映射
	routeMap := make(map[int64]*types.Route)
	var roots []*types.Route
	for _, menu := range menus {
		route := &types.Route{
			ID:   int(menu.MenuId),
			Name: menu.Name,
			Path: menu.Path,
			Meta: types.Meta{
				Title:             menu.Title,
				KeepAlive:         menu.KeepAlive == 1,
				ShowBadge:         menu.ShowBadge == 1,
				IsIframe:          menu.IsIframe == 1,
				IsHide:            menu.IsHide == 1,
				IsHideTab:         menu.IsHideTab == 1,
				IsInMainContainer: menu.IsInMainContainer == 1,
			},
			Children: make([]*types.Route, 0),
		}
		if menu.Component.Valid {
			// 将 RoutesAlias.XXX 格式转换为 ../../viewsRoutesAlias 格式
			component := menu.Component.String
			// if strings.HasPrefix(component, "RoutesAlias.") {
			// 	route.Component = "../../viewsRoutesAlias"
			// } else {
				route.Component = component
			// }
		}
		if menu.Icon.Valid {
			route.Meta.Icon = menu.Icon.String
		}
		if menu.ShowTextBadge.Valid {
			route.Meta.ShowTextBadge = menu.ShowTextBadge.String
		}
		if menu.Link.Valid {
			route.Meta.Link = menu.Link.String
		}
		routeMap[menu.MenuId] = route
	}

	// 2. 组装树结构
	for _, menu := range menus {
		route := routeMap[menu.MenuId]
		if menu.ParentMenuId == 0 {
			roots = append(roots, route)
		} else if parent, ok := routeMap[menu.ParentMenuId]; ok {
			parent.Children = append(parent.Children, route)
		}
	}

	return roots, nil
}

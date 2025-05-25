package menu

import (
	"context"

	"admin-server/api/internal/errorx"
	"admin-server/api/internal/model/user"
	"admin-server/api/internal/pkg/utils"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type MenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuLogic {
	return &MenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuLogic) Menu() (resp *types.MenuInfoResp, err error) {
	userID,err:=utils.GetUserIdFromCtx(l.ctx)
	if err!=nil{
		return nil,err
	}
	userInfo,err:=l.svcCtx.UserModel.FindOne(l.ctx,userID)
	if err!=nil{
		return nil,err
	}
	if userInfo.Status!=user.UserStatusNormal{
       return nil,errorx.ErrUserDisabled
	}
	
	roleUserQuery:=l.svcCtx.RoleUserModel.SelectBuilder().
	Where(squirrel.Eq{
		"user_id":userID,
	})
	roleUserList,err:=l.svcCtx.RoleUserModel.FindAll(l.ctx,roleUserQuery,"")
    if err!=nil{
		return nil,err 
	}
	roleIds:=make([]int64,0,len(roleUserList))
	for _,roleUser:=range roleUserList{
		roleIds=append(roleIds,roleUser.RoleId)
	}
	roleQuery:=l.svcCtx.RoleModel.SelectBuilder().
	Where(squirrel.Eq{"id":roleIds})
	roleList,err:=l.svcCtx.RoleModel.FindAll(l.ctx,roleQuery,"")
	if err!=nil{
		return nil,err
	}

	roleMenuIds:=make([]int64,0,len(roleList))
	for _,v:=range roleList{
		roleMenuIds=append(roleMenuIds,v.Id)
	}
	roleMenuQuery:=l.svcCtx.RoleMenuModel.SelectBuilder().
	Where(squirrel.Eq{"role_id":roleMenuIds})
	roleMenuList,err:=l.svcCtx.RoleMenuModel.FindAll(l.ctx,roleMenuQuery,"")
	if err!=nil{
		return nil,err
	}
	menuIds:=make([]int64,0,len(roleMenuList))
	menuMap:=make(map[int64]struct{},0)
	for _,v:=range roleMenuList{
		if _,ok:=menuMap[v.MenuId];!ok{
			menuMap[v.MenuId]=struct{}{}
			menuIds=append(menuIds,v.MenuId)
		}
	}
	menuQuery:=l.svcCtx.MenuModel.SelectBuilder().
	Where(squirrel.Eq{"id":menuIds})
	menuList,err:=l.svcCtx.MenuModel.FindAll(l.ctx,menuQuery,"")
	if err!=nil{
		return nil,err
	}

	// 构建menuList到Route的映射
	routeMap := make(map[int64]*types.Route)
	var roots []types.Route
	for _, menu := range menuList {
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
	for _, menu := range menuList {
		route := routeMap[menu.MenuId]
		if menu.ParentMenuId == 0 {
			roots = append(roots, *route)
		} else if parent, ok := routeMap[menu.ParentMenuId]; ok {
			parent.Children = append(parent.Children, route)
		}
	}
	
	
	return &types.MenuInfoResp{
		List: roots,
	},nil
}

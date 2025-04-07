package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"context"
	"fmt"
)

var _ MenusModel = (*customMenusModel)(nil)

type (
	// MenusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenusModel.
	MenusModel interface {
		menusModel
		FindByParentId(ctx context.Context, parentId int64) ([]*Menus, error)
		FindByRoleIds(ctx context.Context, roleIds []int64) ([]*Menus, error)
		FindByUserId(ctx context.Context, userId int64) ([]*Menus, error)
	}

	customMenusModel struct {
		*defaultMenusModel
	}
)

// NewMenusModel returns a model for the database table.
func NewMenusModel(conn sqlx.SqlConn, c cache.NodeConf, opts ...cache.Option) MenusModel {
	return &customMenusModel{
		defaultMenusModel: newMenusModel(conn, c, opts...),
	}
}

// FindByParentId 根据父ID查询菜单
func (m *customMenusModel) FindByParentId(ctx context.Context, parentId int64) ([]*Menus, error) {
    query := fmt.Sprintf("select %s from %s where parent_menu_id = ? and is_hide = 0 order by menu_id asc", menuRows, m.table)
    var resp []*Menus
    err := m.QueryRowsNoCacheCtx(ctx, &resp, query, parentId)
    return resp, err
}

func (m *customMenusModel) FindByRoleIds(ctx context.Context, roleIds []int64) ([]*Menus, error) {
	query := fmt.Sprintf("select %s from %s where role_id in (?) and is_hide = 0 order by menu_id asc", menuRows, m.table)
	var resp []*Menus
	err := m.conn.QueryRows(&resp, query, roleIds)
	return resp, err
}

func (m *customMenusModel) FindByUserId(ctx context.Context, userId int64) ([]*Menus, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? and is_hide = 0 order by menu_id asc", menuRows, m.table)
	var resp []*Menus
	err := m.conn.QueryRows(&resp, query, userId)
	return resp, err
}

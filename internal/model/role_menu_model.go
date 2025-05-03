package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"context"
	"fmt"
)

var _ RoleMenuModel = (*customRoleMenuModel)(nil)

type (
	// RoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleMenuModel.
	RoleMenuModel interface {
		roleMenuModel
		FindMenuIDsByRoleId(ctx context.Context,roleId int64)([]*RoleMenu,error)
	}

	customRoleMenuModel struct {
		*defaultRoleMenuModel
	}
)

// NewRoleMenuModel returns a model for the database table.
func NewRoleMenuModel(conn sqlx.SqlConn, c cache.NodeConf, opts ...cache.Option) RoleMenuModel {
	return &customRoleMenuModel{
		defaultRoleMenuModel: newRoleMenuModel(conn, c, opts...),
	}
}

func (m *customRoleMenuModel) FindMenuIDsByRoleId(ctx context.Context,roleId int64)([]*RoleMenu,error){
	query := fmt.Sprintf("select %s from %s where role_id = ?", roleMenuRows, m.table)
	var resp []*RoleMenu
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, roleId)
	return resp, err
}


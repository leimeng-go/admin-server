package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"context"
	"fmt"
)

var _ RoleAuthModel = (*customRoleAuthModel)(nil)

type (
	// RoleAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleAuthModel.
	RoleAuthModel interface {
		roleAuthModel
		FindByRoleId(ctx context.Context,roleId int64)([]*RoleAuth,error)
		FindByAuthId(ctx context.Context,authId int64)([]*RoleAuth,error)
	}

	customRoleAuthModel struct {
		*defaultRoleAuthModel
	}
)

// NewRoleAuthModel returns a model for the database table.
func NewRoleAuthModel(conn sqlx.SqlConn, c cache.NodeConf, opts ...cache.Option) RoleAuthModel {
	return &customRoleAuthModel{
		defaultRoleAuthModel: newRoleAuthModel(conn, c, opts...),
	}
}

func (m *customRoleAuthModel) FindByRoleId(ctx context.Context,roleId int64)([]*RoleAuth,error){
	query := fmt.Sprintf("select %s from %s where role_id = ?", roleAuthRows, m.table)
	var resp []*RoleAuth
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, roleId)
	return resp, err
}

func (m *customRoleAuthModel) FindByAuthId(ctx context.Context,authId int64)([]*RoleAuth,error){
	query := fmt.Sprintf("select %s from %s where auth_id = ?", roleAuthRows, m.table)
	var resp []*RoleAuth
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, authId)
	return resp, err
}
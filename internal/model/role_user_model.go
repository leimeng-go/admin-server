package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"context"
	"fmt"
)

var _ RoleUserModel = (*customRoleUserModel)(nil)

type (
	// RoleUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleUserModel.
	RoleUserModel interface {
		roleUserModel
		FindOneByUserId(ctx context.Context, userID int64) (*RoleUser, error)
	}

	customRoleUserModel struct {
		*defaultRoleUserModel
	}
)

// NewRoleUserModel returns a model for the database table.
func NewRoleUserModel(conn sqlx.SqlConn, c cache.NodeConf, opts ...cache.Option) RoleUserModel {
	return &customRoleUserModel{
		defaultRoleUserModel: newRoleUserModel(conn, c, opts...),
	}
}

func (m *customRoleUserModel) FindOneByUserId(ctx context.Context, userID int64) (*RoleUser, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", roleUserRows, m.table)
	var resp RoleUser
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userID)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
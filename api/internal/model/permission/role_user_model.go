package permission

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleUserModel = (*customRoleUserModel)(nil)

type (
	// RoleUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleUserModel.
	RoleUserModel interface {
		roleUserModel
	}

	customRoleUserModel struct {
		*defaultRoleUserModel
	}
)

// NewRoleUserModel returns a model for the database table.
func NewRoleUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RoleUserModel {
	return &customRoleUserModel{
		defaultRoleUserModel: newRoleUserModel(conn, c, opts...),
	}
}

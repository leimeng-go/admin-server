package permission

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleAuthModel = (*customRoleAuthModel)(nil)

type (
	// RoleAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleAuthModel.
	RoleAuthModel interface {
		roleAuthModel
	}

	customRoleAuthModel struct {
		*defaultRoleAuthModel
	}
)

// NewRoleAuthModel returns a model for the database table.
func NewRoleAuthModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RoleAuthModel {
	return &customRoleAuthModel{
		defaultRoleAuthModel: newRoleAuthModel(conn, c, opts...),
	}
}

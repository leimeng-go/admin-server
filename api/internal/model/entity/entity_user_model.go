package entity

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EntityUserModel = (*customEntityUserModel)(nil)

type (
	// EntityUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEntityUserModel.
	EntityUserModel interface {
		entityUserModel
	}

	customEntityUserModel struct {
		*defaultEntityUserModel
	}
)

// NewEntityUserModel returns a model for the database table.
func NewEntityUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) EntityUserModel {
	return &customEntityUserModel{
		defaultEntityUserModel: newEntityUserModel(conn, c, opts...),
	}
}

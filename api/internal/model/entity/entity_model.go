package entity

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EntityModel = (*customEntityModel)(nil)

type (
	// EntityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEntityModel.
	EntityModel interface {
		entityModel
	}

	customEntityModel struct {
		*defaultEntityModel
	}
)

// NewEntityModel returns a model for the database table.
func NewEntityModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) EntityModel {
	return &customEntityModel{
		defaultEntityModel: newEntityModel(conn, c, opts...),
	}
}

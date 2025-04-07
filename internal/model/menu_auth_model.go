package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MenuAuthModel = (*customMenuAuthModel)(nil)

type (
	// MenuAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenuAuthModel.
	MenuAuthModel interface {
		menuAuthModel
	}

	customMenuAuthModel struct {
		*defaultMenuAuthModel
	}
)

// NewMenuAuthModel returns a model for the database table.
func NewMenuAuthModel(conn sqlx.SqlConn, c cache.NodeConf, opts ...cache.Option) MenuAuthModel {
	return &customMenuAuthModel{
		defaultMenuAuthModel: newMenuAuthModel(conn, c, opts...),
	}
}

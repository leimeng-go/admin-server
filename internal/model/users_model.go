package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)
const (
	UserStatusNormal   = 1
	UserStatusDisabled = 2
	UserStatusDeleted  = 3
)


// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.NodeConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}

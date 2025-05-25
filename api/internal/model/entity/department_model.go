package entity

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DepartmentModel = (*customDepartmentModel)(nil)

type (
	// DepartmentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDepartmentModel.
	DepartmentModel interface {
		departmentModel
	}

	customDepartmentModel struct {
		*defaultDepartmentModel
	}
)
const (
	DepartmentStatusNormal = 1
	DepartmentStatusDisabled = 0
)

// NewDepartmentModel returns a model for the database table.
func NewDepartmentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DepartmentModel {
	return &customDepartmentModel{
		defaultDepartmentModel: newDepartmentModel(conn, c, opts...),
	}
}

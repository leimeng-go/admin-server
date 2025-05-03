package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"fmt"		
)

var _ MenusModel = (*customMenusModel)(nil)

type (
	// MenusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenusModel.
	MenusModel interface {
		menusModel
		FindMenusByIDs(ctx context.Context,ids []int64)([]*Menus,error)		
	}

	customMenusModel struct {
		*defaultMenusModel
	}
)

// NewMenusModel returns a model for the database table.
func NewMenusModel(conn sqlx.SqlConn, c cache.NodeConf, opts ...cache.Option) MenusModel {
	return &customMenusModel{
		defaultMenusModel: newMenusModel(conn, c, opts...),
	}
}

func (m *customMenusModel) FindMenusByIDs(ctx context.Context,ids []int64)([]*Menus,error){
	query := fmt.Sprintf("select %s from %s where id in (?)", menusRows, m.table)
	var resp []*Menus
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, ids)
	return resp, err
}

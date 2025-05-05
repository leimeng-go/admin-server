package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MenuModel = (*customMenuModel)(nil)

type (
	// MenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMenuModel.
	MenuModel interface {
		menuModel
		FindMenusByIDs(ctx context.Context, ids []int64) ([]*Menu, error)

	}

	customMenuModel struct {
		*defaultMenuModel
	}
)

// NewMenuModel returns a model for the database table.
func NewMenuModel(conn sqlx.SqlConn, c cache.NodeConf, opts ...cache.Option) MenuModel {
	return &customMenuModel{
		defaultMenuModel: newMenuModel(conn, c, opts...),
	}
}

func (m *customMenuModel) FindMenusByIDs(ctx context.Context, ids []int64) ([]*Menu, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf("select %s from %s where id in (%s)", menuRows, m.table, strings.Join(placeholders, ","))
	var resp []*Menu
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
	return resp, err
}

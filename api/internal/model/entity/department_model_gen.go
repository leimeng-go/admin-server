// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.8.3

package entity

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"admin-server/api/internal/pkg/globalkey"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	departmentFieldNames          = builder.RawFieldNames(&Department{})
	departmentRows                = strings.Join(departmentFieldNames, ",")
	departmentRowsExpectAutoSet   = strings.Join(stringx.Remove(departmentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	departmentRowsWithPlaceHolder = strings.Join(stringx.Remove(departmentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheDepartmentIdPrefix = "cache:department:id:"
)

type (
	departmentModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Department) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Department, error)
		Update(ctx context.Context, session sqlx.Session, data *Department) (sql.Result, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *Department) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Department, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Department, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Department, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Department, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Department, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultDepartmentModel struct {
		sqlc.CachedConn
		table string
	}

	Department struct {
		Id         int64         `db:"id"`          // 主键
		EntityId   int64         `db:"entity_id"`   // 归属实体ID
		Name       string        `db:"name"`        // 部门名称
		ParentId   sql.NullInt64 `db:"parent_id"`   // 上级部门ID，顶级部门为NULL
		LeaderId   sql.NullInt64 `db:"leader_id"`   // 部门负责人用户ID
		Sort       int64         `db:"sort"`        // 排序
		Status     int64         `db:"status"`      // 状态（1=正常，0=禁用）
		CreateTime time.Time     `db:"create_time"` // 创建时间
		UpdateTime time.Time     `db:"update_time"` // 更新时间
		DeleteTime sql.NullTime  `db:"delete_time"` // 删除时间
		DelState   int64         `db:"del_state"`   // 删除状态 0:未删除 1:已删除
	}
)

func newDepartmentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultDepartmentModel {
	return &defaultDepartmentModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`department`",
	}
}

func (m *defaultDepartmentModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	departmentIdKey := fmt.Sprintf("%s%v", cacheDepartmentIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, departmentIdKey)
	return err
}
func (m *defaultDepartmentModel) FindOne(ctx context.Context, id int64) (*Department, error) {
	departmentIdKey := fmt.Sprintf("%s%v", cacheDepartmentIdPrefix, id)
	var resp Department
	err := m.QueryRowCtx(ctx, &resp, departmentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", departmentRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDepartmentModel) Insert(ctx context.Context, session sqlx.Session, data *Department) (sql.Result, error) {
	data.DeleteTime = sql.NullTime{}
	data.DelState = globalkey.DelStateNo
	departmentIdKey := fmt.Sprintf("%s%v", cacheDepartmentIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, departmentRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.EntityId, data.Name, data.ParentId, data.LeaderId, data.Sort, data.Status, data.DeleteTime, data.DelState)
		}
		return conn.ExecCtx(ctx, query, data.EntityId, data.Name, data.ParentId, data.LeaderId, data.Sort, data.Status, data.DeleteTime, data.DelState)
	}, departmentIdKey)
}

func (m *defaultDepartmentModel) Update(ctx context.Context, session sqlx.Session, data *Department) (sql.Result, error) {
	departmentIdKey := fmt.Sprintf("%s%v", cacheDepartmentIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, departmentRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.EntityId, data.Name, data.ParentId, data.LeaderId, data.Sort, data.Status, data.DeleteTime, data.DelState, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.EntityId, data.Name, data.ParentId, data.LeaderId, data.Sort, data.Status, data.DeleteTime, data.DelState, data.Id)
	}, departmentIdKey)
}

func (m *defaultDepartmentModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *Department) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = sql.NullTime{Time: time.Now(), Valid: true}
	if _, err := m.Update(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "DepartmentModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultDepartmentModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultDepartmentModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultDepartmentModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Department, error) {

	builder = builder.Columns(departmentRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Department
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultDepartmentModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Department, error) {

	builder = builder.Columns(departmentRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Department
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultDepartmentModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Department, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(departmentRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*Department
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultDepartmentModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Department, error) {

	builder = builder.Columns(departmentRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Department
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultDepartmentModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Department, error) {

	builder = builder.Columns(departmentRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Department
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultDepartmentModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultDepartmentModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultDepartmentModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheDepartmentIdPrefix, primary)
}

func (m *defaultDepartmentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", departmentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultDepartmentModel) tableName() string {
	return m.table
}

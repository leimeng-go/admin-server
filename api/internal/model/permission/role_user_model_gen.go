// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.8.3

package permission

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
	roleUserFieldNames          = builder.RawFieldNames(&RoleUser{})
	roleUserRows                = strings.Join(roleUserFieldNames, ",")
	roleUserRowsExpectAutoSet   = strings.Join(stringx.Remove(roleUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	roleUserRowsWithPlaceHolder = strings.Join(stringx.Remove(roleUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheRoleUserIdPrefix           = "cache:roleUser:id:"
	cacheRoleUserRoleIdUserIdPrefix = "cache:roleUser:roleId:userId:"
)

type (
	roleUserModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *RoleUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RoleUser, error)
		FindOneByRoleIdUserId(ctx context.Context, roleId int64, userId int64) (*RoleUser, error)
		Update(ctx context.Context, session sqlx.Session, data *RoleUser) (sql.Result, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *RoleUser) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*RoleUser, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*RoleUser, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*RoleUser, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*RoleUser, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*RoleUser, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultRoleUserModel struct {
		sqlc.CachedConn
		table string
	}

	RoleUser struct {
		Id         int64        `db:"id"`
		RoleId     int64        `db:"role_id"`     // 角色ID
		UserId     int64        `db:"user_id"`     // 用户ID
		CreateTime time.Time    `db:"create_time"` // 创建时间
		UpdateTime time.Time    `db:"update_time"` // 更新时间
		DeleteTime sql.NullTime `db:"delete_time"` // 删除时间
		DelState   int64        `db:"del_state"`   // 删除状态 0:未删除 1:已删除
	}
)

func newRoleUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultRoleUserModel {
	return &defaultRoleUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`role_user`",
	}
}

func (m *defaultRoleUserModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	roleUserIdKey := fmt.Sprintf("%s%v", cacheRoleUserIdPrefix, id)
	roleUserRoleIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheRoleUserRoleIdUserIdPrefix, data.RoleId, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, roleUserIdKey, roleUserRoleIdUserIdKey)
	return err
}
func (m *defaultRoleUserModel) FindOne(ctx context.Context, id int64) (*RoleUser, error) {
	roleUserIdKey := fmt.Sprintf("%s%v", cacheRoleUserIdPrefix, id)
	var resp RoleUser
	err := m.QueryRowCtx(ctx, &resp, roleUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", roleUserRows, m.table)
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

func (m *defaultRoleUserModel) FindOneByRoleIdUserId(ctx context.Context, roleId int64, userId int64) (*RoleUser, error) {
	roleUserRoleIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheRoleUserRoleIdUserIdPrefix, roleId, userId)
	var resp RoleUser
	err := m.QueryRowIndexCtx(ctx, &resp, roleUserRoleIdUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `role_id` = ? and `user_id` = ? and del_state = ? limit 1", roleUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, roleId, userId, globalkey.DelStateNo); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRoleUserModel) Insert(ctx context.Context, session sqlx.Session, data *RoleUser) (sql.Result, error) {
	data.DeleteTime = sql.NullTime{}
	data.DelState = globalkey.DelStateNo
	roleUserIdKey := fmt.Sprintf("%s%v", cacheRoleUserIdPrefix, data.Id)
	roleUserRoleIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheRoleUserRoleIdUserIdPrefix, data.RoleId, data.UserId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, roleUserRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.RoleId, data.UserId, data.DeleteTime, data.DelState)
		}
		return conn.ExecCtx(ctx, query, data.RoleId, data.UserId, data.DeleteTime, data.DelState)
	}, roleUserIdKey, roleUserRoleIdUserIdKey)
}

func (m *defaultRoleUserModel) Update(ctx context.Context, session sqlx.Session, newData *RoleUser) (sql.Result, error) {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return nil, err
	}
	roleUserIdKey := fmt.Sprintf("%s%v", cacheRoleUserIdPrefix, data.Id)
	roleUserRoleIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheRoleUserRoleIdUserIdPrefix, data.RoleId, data.UserId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, roleUserRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.RoleId, newData.UserId, newData.DeleteTime, newData.DelState, newData.Id)
		}
		return conn.ExecCtx(ctx, query, newData.RoleId, newData.UserId, newData.DeleteTime, newData.DelState, newData.Id)
	}, roleUserIdKey, roleUserRoleIdUserIdKey)
}

func (m *defaultRoleUserModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *RoleUser) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = sql.NullTime{Time: time.Now(), Valid: true}
	if _, err := m.Update(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "RoleUserModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultRoleUserModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultRoleUserModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultRoleUserModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*RoleUser, error) {

	builder = builder.Columns(roleUserRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*RoleUser
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRoleUserModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*RoleUser, error) {

	builder = builder.Columns(roleUserRows)

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

	var resp []*RoleUser
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRoleUserModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*RoleUser, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(roleUserRows)

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

	var resp []*RoleUser
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultRoleUserModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*RoleUser, error) {

	builder = builder.Columns(roleUserRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*RoleUser
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRoleUserModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*RoleUser, error) {

	builder = builder.Columns(roleUserRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*RoleUser
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRoleUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultRoleUserModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultRoleUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheRoleUserIdPrefix, primary)
}

func (m *defaultRoleUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", roleUserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRoleUserModel) tableName() string {
	return m.table
}

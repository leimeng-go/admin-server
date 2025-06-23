package entity

import (
	"admin-server/api/internal/pkg/globalkey"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	departmentTreeFieldNames          = builder.RawFieldNames(&DepartmentTree{})
	departmentTreeRows                = strings.Join(departmentTreeFieldNames, ",")
	departmentTreeRowsExpectAutoSet   = strings.Join(stringx.Remove(departmentTreeFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	departmentTreeRowsWithPlaceHolder = strings.Join(stringx.Remove(departmentTreeFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheDepartmentTreeIdPrefix   = "cache:department_tree:id:"
	cacheDepartmentTreeCodePrefix = "cache:department_tree:code:"
)

type (
	DepartmentTreeModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *DepartmentTree) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DepartmentTree, error)
		FindOneByCode(ctx context.Context, entityId int64, code string) (*DepartmentTree, error)
		Update(ctx context.Context, session sqlx.Session, data *DepartmentTree) (sql.Result, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
		DeleteSoft(ctx context.Context, session sqlx.Session, data *DepartmentTree) error

		// 树形结构相关方法
		FindChildren(ctx context.Context, entityId, parentId int64) ([]*DepartmentTree, error)
		FindAllChildren(ctx context.Context, entityId, departmentId int64, includeSelf bool) ([]*DepartmentTree, error)
		FindParents(ctx context.Context, entityId, departmentId int64, includeSelf bool) ([]*DepartmentTree, error)
		FindByLevel(ctx context.Context, entityId int64, level int) ([]*DepartmentTree, error)
		FindTree(ctx context.Context, entityId int64) ([]*DepartmentTree, error)
		FindSiblings(ctx context.Context, entityId, departmentId int64, includeSelf bool) ([]*DepartmentTree, error)
		MoveDepartment(ctx context.Context, session sqlx.Session, departmentId, newParentId int64) error
	}

	defaultDepartmentTreeModel struct {
		sqlc.CachedConn
		table string
	}

	DepartmentTree struct {
		Id          int64          `db:"id"`          // 主键
		EntityId    int64          `db:"entity_id"`   // 归属实体ID
		Name        string         `db:"name"`        // 部门名称
		Code        sql.NullString `db:"code"`        // 部门编码
		ParentId    sql.NullInt64  `db:"parent_id"`   // 上级部门ID
		Ancestors   sql.NullString `db:"ancestors"`   // 祖级路径
		Level       int64          `db:"level"`       // 层级深度
		Path        sql.NullString `db:"path"`        // 完整路径
		LeftValue   sql.NullInt64  `db:"left_value"`  // 左值
		RightValue  sql.NullInt64  `db:"right_value"` // 右值
		LeaderId    sql.NullInt64  `db:"leader_id"`   // 部门负责人用户ID
		LeaderName  sql.NullString `db:"leader_name"` // 部门负责人姓名
		Phone       sql.NullString `db:"phone"`       // 部门联系电话
		Email       sql.NullString `db:"email"`       // 部门邮箱
		Description sql.NullString `db:"description"` // 部门描述
		Sort        int64          `db:"sort"`        // 排序
		Status      int64          `db:"status"`      // 状态
		IsLeaf      int64          `db:"is_leaf"`     // 是否叶子节点
		CreateTime  sql.NullTime   `db:"create_time"` // 创建时间
		UpdateTime  sql.NullTime   `db:"update_time"` // 更新时间
		DeleteTime  sql.NullTime   `db:"delete_time"` // 删除时间
		DelState    int64          `db:"del_state"`   // 删除状态
	}
)

func NewDepartmentTreeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DepartmentTreeModel {
	return &defaultDepartmentTreeModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`department_tree`",
	}
}

func (m *defaultDepartmentTreeModel) Insert(ctx context.Context, session sqlx.Session, data *DepartmentTree) (sql.Result, error) {
	data.DelState = globalkey.DelStateNo
	departmentTreeIdKey := fmt.Sprintf("%s%v", cacheDepartmentTreeIdPrefix, data.Id)
	departmentTreeCodeKey := fmt.Sprintf("%s%v", cacheDepartmentTreeCodePrefix, data.Code.String)

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			m.table, departmentTreeRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query,
				data.EntityId, data.Name, data.Code, data.ParentId, data.Ancestors,
				data.Level, data.Path, data.LeftValue, data.RightValue, data.LeaderId,
				data.LeaderName, data.Phone, data.Email, data.Description, data.Sort,
				data.Status, data.IsLeaf, data.DeleteTime, data.DelState)
		}
		return conn.ExecCtx(ctx, query,
			data.EntityId, data.Name, data.Code, data.ParentId, data.Ancestors,
			data.Level, data.Path, data.LeftValue, data.RightValue, data.LeaderId,
			data.LeaderName, data.Phone, data.Email, data.Description, data.Sort,
			data.Status, data.IsLeaf, data.DeleteTime, data.DelState)
	}, departmentTreeIdKey, departmentTreeCodeKey)
}

func (m *defaultDepartmentTreeModel) FindOne(ctx context.Context, id int64) (*DepartmentTree, error) {
	departmentTreeIdKey := fmt.Sprintf("%s%v", cacheDepartmentTreeIdPrefix, id)
	var resp DepartmentTree
	err := m.QueryRowCtx(ctx, &resp, departmentTreeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", departmentTreeRows, m.table)
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

func (m *defaultDepartmentTreeModel) FindOneByCode(ctx context.Context, entityId int64, code string) (*DepartmentTree, error) {
	departmentTreeCodeKey := fmt.Sprintf("%s%v", cacheDepartmentTreeCodePrefix, code)
	var resp DepartmentTree
	err := m.QueryRowIndexCtx(ctx, &resp, departmentTreeCodeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `entity_id` = ? and `code` = ? and del_state = ? limit 1", departmentTreeRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, entityId, code, globalkey.DelStateNo); err != nil {
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

func (m *defaultDepartmentTreeModel) Update(ctx context.Context, session sqlx.Session, data *DepartmentTree) (sql.Result, error) {
	departmentTreeIdKey := fmt.Sprintf("%s%v", cacheDepartmentTreeIdPrefix, data.Id)
	departmentTreeCodeKey := fmt.Sprintf("%s%v", cacheDepartmentTreeCodePrefix, data.Code.String)

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, departmentTreeRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query,
				data.EntityId, data.Name, data.Code, data.ParentId, data.Ancestors,
				data.Level, data.Path, data.LeftValue, data.RightValue, data.LeaderId,
				data.LeaderName, data.Phone, data.Email, data.Description, data.Sort,
				data.Status, data.IsLeaf, data.DeleteTime, data.DelState, data.Id)
		}
		return conn.ExecCtx(ctx, query,
			data.EntityId, data.Name, data.Code, data.ParentId, data.Ancestors,
			data.Level, data.Path, data.LeftValue, data.RightValue, data.LeaderId,
			data.LeaderName, data.Phone, data.Email, data.Description, data.Sort,
			data.Status, data.IsLeaf, data.DeleteTime, data.DelState, data.Id)
	}, departmentTreeIdKey, departmentTreeCodeKey)
}

func (m *defaultDepartmentTreeModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	departmentTreeIdKey := fmt.Sprintf("%s%v", cacheDepartmentTreeIdPrefix, id)
	departmentTreeCodeKey := fmt.Sprintf("%s%v", cacheDepartmentTreeCodePrefix, data.Code.String)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, departmentTreeIdKey, departmentTreeCodeKey)
	return err
}

func (m *defaultDepartmentTreeModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *DepartmentTree) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = sql.NullTime{Time: time.Now(), Valid: true}
	if _, err := m.Update(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "DepartmentTreeModel delete err : %+v", err)
	}
	return nil
}

// 查找直接子部门
func (m *defaultDepartmentTreeModel) FindChildren(ctx context.Context, entityId, parentId int64) ([]*DepartmentTree, error) {
	query := fmt.Sprintf("select %s from %s where `entity_id` = ? and `parent_id` = ? and del_state = ? order by `sort`, `id`",
		departmentTreeRows, m.table)
	var resp []*DepartmentTree
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, entityId, parentId, globalkey.DelStateNo)
	return resp, err
}

// 查找所有子部门（包括子部门的子部门）
func (m *defaultDepartmentTreeModel) FindAllChildren(ctx context.Context, entityId, departmentId int64, includeSelf bool) ([]*DepartmentTree, error) {
	// 使用存储过程
	query := "CALL sp_get_department_children(?, ?, ?)"
	var resp []*DepartmentTree
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, entityId, departmentId, includeSelf)
	return resp, err
}

// 查找所有父部门
func (m *defaultDepartmentTreeModel) FindParents(ctx context.Context, entityId, departmentId int64, includeSelf bool) ([]*DepartmentTree, error) {
	// 使用存储过程
	query := "CALL sp_get_department_parents(?, ?, ?)"
	var resp []*DepartmentTree
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, entityId, departmentId, includeSelf)
	return resp, err
}

// 根据层级查找部门
func (m *defaultDepartmentTreeModel) FindByLevel(ctx context.Context, entityId int64, level int) ([]*DepartmentTree, error) {
	query := fmt.Sprintf("select %s from %s where `entity_id` = ? and `level` = ? and del_state = ? order by `sort`, `id`",
		departmentTreeRows, m.table)
	var resp []*DepartmentTree
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, entityId, level, globalkey.DelStateNo)
	return resp, err
}

// 查找完整的部门树
func (m *defaultDepartmentTreeModel) FindTree(ctx context.Context, entityId int64) ([]*DepartmentTree, error) {
	query := fmt.Sprintf("select %s from %s where `entity_id` = ? and del_state = ? order by `left_value`",
		departmentTreeRows, m.table)
	var resp []*DepartmentTree
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, entityId, globalkey.DelStateNo)
	return resp, err
}

// 查找兄弟部门
func (m *defaultDepartmentTreeModel) FindSiblings(ctx context.Context, entityId, departmentId int64, includeSelf bool) ([]*DepartmentTree, error) {
	// 先获取当前部门的父部门ID
	dept, err := m.FindOne(ctx, departmentId)
	if err != nil {
		return nil, err
	}

	var parentId sql.NullInt64
	if dept.ParentId.Valid {
		parentId = dept.ParentId
	}

	query := fmt.Sprintf("select %s from %s where `entity_id` = ? and `parent_id` %s and del_state = ? order by `sort`, `id`",
		departmentTreeRows, m.table,
		func() string {
			if parentId.Valid {
				return "= ?"
			}
			return "IS NULL"
		}())

	var resp []*DepartmentTree
	if parentId.Valid {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, entityId, parentId.Int64, globalkey.DelStateNo)
	} else {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, entityId, globalkey.DelStateNo)
	}

	if err != nil {
		return nil, err
	}

	// 如果不包含自己，过滤掉当前部门
	if !includeSelf {
		var filtered []*DepartmentTree
		for _, item := range resp {
			if item.Id != departmentId {
				filtered = append(filtered, item)
			}
		}
		return filtered, nil
	}

	return resp, nil
}

// 移动部门（改变父部门）
func (m *defaultDepartmentTreeModel) MoveDepartment(ctx context.Context, session sqlx.Session, departmentId, newParentId int64) error {
	// 这是一个复杂的操作，需要重新计算整个子树的结构
	// 建议通过应用程序逻辑来处理，或者使用专门的存储过程
	return errors.New("MoveDepartment operation should be implemented with application logic")
}

func (m *defaultDepartmentTreeModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheDepartmentTreeIdPrefix, primary)
}

func (m *defaultDepartmentTreeModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", departmentTreeRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

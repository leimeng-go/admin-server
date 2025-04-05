package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	userFieldNames        = strings.Join([]string{"id", "username", "password", "nickname", "avatar", "email", "mobile", "role", "status", "created_at", "updated_at", "deleted_at"}, ",")
	userRows              = strings.Join([]string{"?", "?", "?", "?", "?", "?", "?", "?", "?", "CURRENT_TIMESTAMP", "CURRENT_TIMESTAMP", "?"}, ",")
	cacheUserIdPrefix     = "cache:user:id:"
	cacheUserNamePrefix   = "cache:user:name:"
	cacheUserEmailPrefix  = "cache:user:email:"
	cacheUserMobilePrefix = "cache:user:mobile:"
)

type (
	User struct {
		Id        uint64       `db:"id"`
		Username  string       `db:"username"`
		Password  string       `db:"password"`
		Nickname  string       `db:"nickname"`
		Avatar    string       `db:"avatar"`
		Email     string       `db:"email"`
		Mobile    string       `db:"mobile"`
		Role      string       `db:"role"`
		Status    int64        `db:"status"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt time.Time    `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
	}

	UserModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*User, error)
		FindOneByUsername(ctx context.Context, username string) (*User, error)
		FindOneByEmail(ctx context.Context, email string) (*User, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id uint64) error
		List(ctx context.Context, page, pageSize int64, keyword string) ([]*User, int64, error)
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}
)

func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "users",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", m.table, userFieldNames, userRows)
	return m.conn.ExecCtx(ctx, query, data.Id, data.Username, data.Password, data.Nickname,
		data.Avatar, data.Email, data.Mobile, data.Role, data.Status, data.DeletedAt)
}

func (m *defaultUserModel) FindOne(ctx context.Context, id uint64) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ? AND deleted_at IS NULL LIMIT 1", userFieldNames, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE username = ? AND deleted_at IS NULL LIMIT 1", userFieldNames, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE email = ? AND deleted_at IS NULL LIMIT 1", userFieldNames, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE mobile = ? AND deleted_at IS NULL LIMIT 1", userFieldNames, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, mobile)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	query := fmt.Sprintf("UPDATE %s SET nickname=?, avatar=?, email=?, mobile=?, role=?, status=?, updated_at=CURRENT_TIMESTAMP WHERE id=? AND deleted_at IS NULL", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.Nickname, data.Avatar, data.Email, data.Mobile, data.Role, data.Status, data.Id)
	return err
}

func (m *defaultUserModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=CURRENT_TIMESTAMP, updated_at=CURRENT_TIMESTAMP WHERE id=? AND deleted_at IS NULL", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserModel) List(ctx context.Context, page, pageSize int64, keyword string) ([]*User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	where := "WHERE deleted_at IS NULL"
	args := make([]interface{}, 0)
	if keyword != "" {
		where += " AND (username LIKE ? OR nickname LIKE ? OR email LIKE ? OR mobile LIKE ?)"
		keyword = "%" + keyword + "%"
		args = append(args, keyword, keyword, keyword, keyword)
	}

	// 查询总数
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s %s", m.table, where)
	var total int64
	err := m.conn.QueryRowCtx(ctx, &total, countQuery, args...)
	if err != nil {
		return nil, 0, err
	}

	// 查询数据
	query := fmt.Sprintf("SELECT %s FROM %s %s ORDER BY id DESC LIMIT ? OFFSET ?", userFieldNames, m.table, where)
	args = append(args, pageSize, offset)
	var resp []*User
	err = m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

const (
	UserStatusNormal   = 1
	UserStatusDisabled = 2
	UserStatusDeleted  = 3

	createUserTable = `CREATE TABLE IF NOT EXISTS users (
		id BIGINT UNSIGNED PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL,
		nickname VARCHAR(50) NOT NULL,
		avatar VARCHAR(255) DEFAULT '',
		email VARCHAR(100) NOT NULL DEFAULT '',
		mobile VARCHAR(20) NOT NULL DEFAULT '',
		role VARCHAR(20) NOT NULL DEFAULT 'user',
		status TINYINT NOT NULL DEFAULT 1,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		INDEX idx_username (username),
		INDEX idx_email (email)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
)

func InitTables(db *sql.DB) error {
	_, err := db.Exec(createUserTable)
	return err
}

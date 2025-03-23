package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type User struct {
	Id        int64     `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Nickname  string    `db:"nickname"`
	Avatar    string    `db:"avatar"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	Status    int       `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserModel interface {
	Insert(ctx context.Context, data *User) (sql.Result, error)
	FindOne(ctx context.Context, id int64) (*User, error)
	FindOneByUsername(ctx context.Context, username string) (*User, error)
	FindOneByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, data *User) error
	Delete(ctx context.Context, id int64) error
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindList(ctx context.Context, page, pageSize int64, keyword string) ([]*User, int64, error)
	SoftDelete(ctx context.Context, id int64) error
}

type defaultUserModel struct {
	conn  sqlx.SqlConn
	table string
}

func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "users",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (username, password, nickname, avatar, email, role, status) values (?, ?, ?, ?, ?, ?, ?)", m.table)
	return m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Nickname, data.Avatar, data.Email, data.Role, data.Status)
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	query := fmt.Sprintf("select * from %s where id = ? limit 1", m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	query := fmt.Sprintf("select * from %s where username = ? limit 1", m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	query := fmt.Sprintf("select * from %s where email = ? limit 1", m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	query := fmt.Sprintf("update %s set nickname = ?, avatar = ?, email = ?, role = ?, status = ? where id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.Nickname, data.Avatar, data.Email, data.Role, data.Status, data.Id)
	return err
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("update %s set status = ? where id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, UserStatusDeleted, id)
	return err
}

func (m *defaultUserModel) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	query := fmt.Sprintf("select count(*) from %s where username = ?", m.table)
	var count int
	err := m.conn.QueryRowCtx(ctx, &count, query, username)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *defaultUserModel) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	query := fmt.Sprintf("select count(*) from %s where email = ?", m.table)
	var count int
	err := m.conn.QueryRowCtx(ctx, &count, query, email)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *defaultUserModel) FindList(ctx context.Context, page, pageSize int64, keyword string) ([]*User, int64, error) {
	var conditions []string
	var args []interface{}
	if keyword != "" {
		conditions = append(conditions, "(username like ? or nickname like ?)")
		args = append(args, "%"+keyword+"%", "%"+keyword+"%")
	}

	where := "where 1=1"
	if len(conditions) > 0 {
		where += " and " + strings.Join(conditions, " and ")
	}

	countQuery := fmt.Sprintf("select count(*) from %s %s", m.table, where)
	var total int64
	err := m.conn.QueryRowCtx(ctx, &total, countQuery, args...)
	if err != nil {
		return nil, 0, err
	}

	query := fmt.Sprintf("select * from %s %s limit ?,?", m.table, where)
	args = append(args, (page-1)*pageSize, pageSize)
	var resp []*User
	err = m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func (m *defaultUserModel) SoftDelete(ctx context.Context, id int64) error {
	return m.Delete(ctx, id)
}

const (
	UserStatusNormal   = 1
	UserStatusDisabled = 2
	UserStatusDeleted  = 3

	createUserTable = `CREATE TABLE IF NOT EXISTS users (
		id BIGINT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(100) NOT NULL,
		nickname VARCHAR(50) NOT NULL,
		avatar VARCHAR(255) DEFAULT '',
		email VARCHAR(100) NOT NULL DEFAULT '',
		role VARCHAR(20) NOT NULL DEFAULT 'user',
		status TINYINT NOT NULL DEFAULT 1,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		INDEX idx_username (username),
		INDEX idx_email (email)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
)

func InitTables(db *sql.DB) error {
	_, err := db.Exec(createUserTable)
	return err
}

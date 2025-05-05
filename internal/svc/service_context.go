package svc

import (
	"admin-server/internal/config"
	"admin-server/internal/model"
	"admin-server/internal/utils"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UsersModel
	RoleModel      model.RoleModel
	UserRoleModel  model.RoleUserModel
	RoleMenuModel  model.RoleMenuModel
	MenuModel      model.MenuModel
	Snowflake      *utils.Snowflake
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	snowflake, _ := utils.NewSnowflake(1) // 使用机器ID 1
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUsersModel(conn, c.CacheRedis),
		RoleModel:      model.NewRoleModel(conn, c.CacheRedis),
		UserRoleModel:  model.NewRoleUserModel(conn, c.CacheRedis),
		RoleMenuModel:  model.NewRoleMenuModel(conn, c.CacheRedis),
		MenuModel:      model.NewMenuModel(conn, c.CacheRedis),
		Snowflake:      snowflake,
		// AuthMiddleware: rest.NewJwtAuthMiddleware(c.Auth.AccessSecret).Handle,
	}
}

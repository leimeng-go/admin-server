package svc

import (
	"admin-server/internal/config"
	"admin-server/internal/model"
	"admin-server/internal/utils"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel
	MenusModel     model.MenusModel
	RoleModel      model.RoleModel
	UserRoleModel  model.RoleUserModel
	RoleMenuModel  model.RoleMenuModel
	MenuModel      model.MenusModel
	Snowflake      *utils.Snowflake
	// AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	snowflake, _ := utils.NewSnowflake(1) // 使用机器ID 1
	// rest.WithUnauthorizedCallback()
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(conn),
		Snowflake:      snowflake,
		// AuthMiddleware: rest.NewJwtAuthMiddleware(c.Auth.AccessSecret).Handle,
	}
}

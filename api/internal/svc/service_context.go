package svc

import (
	"admin-server/api/internal/config"
	"admin-server/api/internal/model/permission"
	"admin-server/api/internal/model/route"
	"admin-server/api/internal/model/user"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserModel user.UserModel
	RoleModel permission.RoleModel
	RoleAuthModel permission.RoleAuthModel
	RoleMenuModel permission.RoleMenuModel
	RoleUserModel permission.RoleUserModel
	MenuModel route.MenuModel
	MenuAuthModel route.MenuAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn:=sqlx.NewMysql(c.Database.Source)
	return &ServiceContext{
		Config: c,
		UserModel: user.NewUserModel(conn,cache.ClusterConf{c.CacheRedis}),
	}
}

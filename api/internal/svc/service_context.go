package svc

import (
	"admin-server/api/internal/config"
	"admin-server/api/internal/model/entity"
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
	EntityModel entity.EntityModel
	EntityUserModel entity.EntityUserModel
	DepartmentModel entity.DepartmentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn:=sqlx.NewMysql(c.Database.Source)
	return &ServiceContext{
		Config: c,
		UserModel: user.NewUserModel(conn,cache.ClusterConf{c.CacheRedis}),
		RoleModel: permission.NewRoleModel(conn,cache.ClusterConf{c.CacheRedis}),
		RoleAuthModel: permission.NewRoleAuthModel(conn,cache.ClusterConf{c.CacheRedis}),
		RoleMenuModel: permission.NewRoleMenuModel(conn,cache.ClusterConf{c.CacheRedis}),
		RoleUserModel: permission.NewRoleUserModel(conn,cache.ClusterConf{c.CacheRedis}),
		MenuModel: route.NewMenuModel(conn,cache.ClusterConf{c.CacheRedis}),
		MenuAuthModel: route.NewMenuAuthModel(conn,cache.ClusterConf{c.CacheRedis}),
		EntityModel: entity.NewEntityModel(conn,cache.ClusterConf{c.CacheRedis}),
		EntityUserModel: entity.NewEntityUserModel(conn,cache.ClusterConf{c.CacheRedis}),
		DepartmentModel: entity.NewDepartmentModel(conn,cache.ClusterConf{c.CacheRedis}),
	}
}

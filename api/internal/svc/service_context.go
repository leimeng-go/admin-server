package svc

import (
	"admin-server/api/internal/config"
	"admin-server/api/internal/model/entity"
	"admin-server/api/internal/model/permission"
	"admin-server/api/internal/model/route"
	"admin-server/api/internal/model/user"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	Enforcer        *casbin.Enforcer
	UserModel       user.UserModel
	RoleModel       permission.RoleModel
	RoleAuthModel   permission.RoleAuthModel
	RoleMenuModel   permission.RoleMenuModel
	RoleUserModel   permission.RoleUserModel
	MenuModel       route.MenuModel
	MenuAuthModel   route.MenuAuthModel
	EntityModel     entity.EntityModel
	EntityUserModel entity.EntityUserModel
	DepartmentModel entity.DepartmentModel
}


func NewServiceContext(c config.Config) *ServiceContext {
	db,err:=gorm.Open(mysql.Open(c.Casbin.Source),&gorm.Config{})
	if err!=nil{
		logx.Must(err)
	}
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		logx.Must(err)
	}
	m, err := model.NewModelFromString(c.RBACModel)
	if err != nil {
		logx.Must(err)
	}
	enforcer, err := casbin.NewEnforcer(m, a)
	if err != nil {
		logx.Must(err)
	}
	enforcer.LoadPolicy()
	conn := sqlx.NewMysql(c.Database.Source)
	return &ServiceContext{
		Config:          c,
		Enforcer:        enforcer,
		UserModel:       user.NewUserModel(conn, cache.ClusterConf{c.CacheRedis}),
		RoleModel:       permission.NewRoleModel(conn, cache.ClusterConf{c.CacheRedis}),
		RoleAuthModel:   permission.NewRoleAuthModel(conn, cache.ClusterConf{c.CacheRedis}),
		RoleMenuModel:   permission.NewRoleMenuModel(conn, cache.ClusterConf{c.CacheRedis}),
		RoleUserModel:   permission.NewRoleUserModel(conn, cache.ClusterConf{c.CacheRedis}),
		MenuModel:       route.NewMenuModel(conn, cache.ClusterConf{c.CacheRedis}),
		MenuAuthModel:   route.NewMenuAuthModel(conn, cache.ClusterConf{c.CacheRedis}),
		EntityModel:     entity.NewEntityModel(conn, cache.ClusterConf{c.CacheRedis}),
		EntityUserModel: entity.NewEntityUserModel(conn, cache.ClusterConf{c.CacheRedis}),
		DepartmentModel: entity.NewDepartmentModel(conn, cache.ClusterConf{c.CacheRedis}),
	}
}

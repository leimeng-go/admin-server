package svc

import (
	"github.com/leimeng-go/admin-server/internal/config"
	"github.com/leimeng-go/admin-server/internal/middleware"
	"github.com/leimeng-go/admin-server/internal/model"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	UserModel      model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		UserModel:      model.NewUserModel(conn),
	}
}

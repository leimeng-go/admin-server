package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	CacheRedis cache.NodeConf
	Database   struct {
		Driver string
		Source string
	}
	Redis struct {
		Host string
		Type string
		Pass string
	}

	Email struct {
		Host     string
		Port     int
		Username string
		Password string
		From     string
		FromName string
	}
	RBACModel string
	Casbin struct {
		Driver string
		Source string
	}
}

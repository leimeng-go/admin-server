package main

import (
	"admin-server/api/internal/config"
	"admin-server/api/internal/errorx"
	"admin-server/api/internal/handler"
	"admin-server/api/internal/svc"
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

    httpx.SetErrorHandler(func(err error) (int, any) {
       if e,ok:=err.(*errorx.Error);ok{
		httpCode:=e.GetHTTPCode()
        return httpCode,map[string]any{"code":e.Code,"message":e.Message}
       }
       return http.StatusInternalServerError,map[string]any{"code":500,"message":"服务器内部错误"}
    })
	httpx.SetOkHandler(func(ctx context.Context, v any) any {
		return map[string]any{"code": errorx.Success, "message": "success", "data": v}
	})
	server := rest.MustNewServer(c.RestConf,rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

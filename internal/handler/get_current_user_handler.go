package handler

import (
	"net/http"

	"github.com/leimeng-go/admin-server/internal/logic"
	"github.com/leimeng-go/admin-server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取当前登录用户的详细信息
func getCurrentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetCurrentUserLogic(r.Context(), svcCtx)
		resp, err := l.GetCurrentUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

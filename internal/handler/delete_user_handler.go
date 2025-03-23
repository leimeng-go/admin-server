package handler

import (
	"net/http"

	"github.com/leimeng-go/admin-server/internal/logic"
	"github.com/leimeng-go/admin-server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除指定用户（软删除）
func deleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDeleteUserLogic(r.Context(), svcCtx)
		err := l.DeleteUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

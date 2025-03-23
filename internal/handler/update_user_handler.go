package handler

import (
	"net/http"

	"github.com/leimeng-go/admin-server/internal/logic"
	"github.com/leimeng-go/admin-server/internal/svc"
	"github.com/leimeng-go/admin-server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新指定用户的昵称、头像或角色
func updateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateUserLogic(r.Context(), svcCtx)
		err := l.UpdateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

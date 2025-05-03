package user

import (
	"net/http"

	"admin-server/internal/logic/user"
	"admin-server/internal/svc"
	"admin-server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 管理员创建新用户，可以指定用户角色
func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewCreateUserLogic(r.Context(), svcCtx)
		err := l.CreateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

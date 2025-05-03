package user

import (
	"net/http"

	"admin-server/internal/errorx"
	"admin-server/internal/logic/user"
	"admin-server/internal/svc"
	"admin-server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 使用用户名和密码登录系统
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			errorx.WriteError(w, errorx.ErrInvalidParams)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			errorx.WriteError(w, err)
			return
		}

		errorx.WriteSuccess(w, resp)
	}
}

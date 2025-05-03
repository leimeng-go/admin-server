package user

import (
	"net/http"

	"admin-server/internal/logic/user"
	"admin-server/internal/svc"
	"admin-server/internal/errorx"
)

// 获取当前登录用户的详细信息
func GetCurrentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetCurrentUserLogic(r.Context(), svcCtx)
		resp, err := l.GetCurrentUser()
		if err != nil {
			errorx.WriteError(w, err)
			return
		} else {
			errorx.WriteSuccess(w, resp)
		}
	}
}

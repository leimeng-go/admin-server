package user

import (
	"net/http"

	"admin-server/internal/logic/user"
	"admin-server/internal/svc"
	"admin-server/internal/errorx"
)

func GetUserMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserMenuLogic(r.Context(), svcCtx)
		resp, err := l.GetUserMenu()
		if err != nil {
			errorx.WriteError(w, err)
			return
		} else {
			errorx.WriteSuccess(w, resp)
		}
	}
}

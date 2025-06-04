package role

import (
	"net/http"

	"admin-server/api/internal/logic/role"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddroleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewAddroleLogic(r.Context(), svcCtx)
		err := l.Addrole(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}

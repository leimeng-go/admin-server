package entity

import (
	"net/http"

	"admin-server/api/internal/logic/entity"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AdddepartmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddDepartmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := entity.NewAdddepartmentLogic(r.Context(), svcCtx)
		err := l.Adddepartment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

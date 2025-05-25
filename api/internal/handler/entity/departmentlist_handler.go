package entity

import (
	"net/http"

	"admin-server/api/internal/logic/entity"
	"admin-server/api/internal/svc"
	"admin-server/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DepartmentlistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DepartmentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := entity.NewDepartmentlistLogic(r.Context(), svcCtx)
		resp, err := l.Departmentlist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package entity

import (
	"net/http"

	"admin-server/api/internal/logic/entity"
	"admin-server/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdatedepartmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := entity.NewUpdatedepartmentLogic(r.Context(), svcCtx)
		resp, err := l.Updatedepartment()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package user

import (
	"net/http"

	"admin-server/internal/logic/user"
	"admin-server/internal/svc"
	"admin-server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 分页获取用户列表，支持关键词搜索
func GetUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

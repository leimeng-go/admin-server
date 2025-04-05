package handler

// import (
// 	"net/http"

// 	"admin-server/internal/logic"
// 	"admin-server/internal/svc"
// 	"admin-server/internal/types"

// 	"github.com/zeromicro/go-zero/rest/httpx"
// )

// func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var req types.RegisterReq
// 		if err := httpx.Parse(r, &req); err != nil {
// 			httpx.Error(w, err)
// 			return
// 		}

// 		l := logic.NewUserLogic(r.Context(), svcCtx)
// 		err := l.Register(&req)
// 		if err != nil {
// 			httpx.Error(w, err)
// 		} else {
// 			httpx.Ok(w)
// 		}
// 	}
// }

// func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var req types.LoginReq
// 		if err := httpx.Parse(r, &req); err != nil {
// 			httpx.Error(w, err)
// 			return
// 		}

// 		l := logic.NewUserLogic(r.Context(), svcCtx)
// 		resp, err := l.Login(&req)
// 		if err != nil {
// 			httpx.Error(w, err)
// 		} else {
// 			httpx.OkJson(w, resp)
// 		}
// 	}
// }

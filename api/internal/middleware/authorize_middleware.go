package middleware

import (
	"net/http"
    "admin-server/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/handler"
)

type AuthorizeMiddleware struct {
	svcCtx *svc.ServiceContext
}

func NewAuthorizeMiddleware(svcCtx *svc.ServiceContext) *AuthorizeMiddleware {
	return &AuthorizeMiddleware{
		svcCtx: svcCtx,
	}
}

func (m *AuthorizeMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		handler.Authorize(m.svcCtx.Config.Auth.AccessSecret,handler.WithUnauthorizedCallback())

		// Passthrough to next handler if need
		next(w, r)
	}
}

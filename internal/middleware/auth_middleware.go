package middleware

import (
	"admin-server/internal/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	accessSecret string
}

func NewAuthMiddleware(accessSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		accessSecret: accessSecret,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			httpx.OkJson(w, map[string]interface{}{
				"code": http.StatusUnauthorized,
				"msg":  "missing authorization header",
			})
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			httpx.OkJson(w, map[string]interface{}{
				"code": http.StatusUnauthorized,
				"msg":  "invalid authorization header",
			})
			return
		}

		claims, err := utils.ParseToken(parts[1], m.accessSecret)
		if err != nil {
			logx.Errorf("Failed to parse token: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			httpx.OkJson(w, map[string]interface{}{
				"code": http.StatusUnauthorized,
				"msg":  "invalid token",
			})
			return
		}

		r.Header.Set("X-User-ID", strconv.FormatInt(claims.UserId, 10))
		next(w, r)
	}
}

func (m *AuthMiddleware) GenerateToken(userId string, role string) (string, error) {
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return "", err
	}
	return utils.GenerateToken(id, role, m.accessSecret)
}

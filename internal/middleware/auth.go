package middleware

// import (
// 	"net/http"

// 	"admin-server/internal/errorx"
// 	"admin-server/internal/svc"
// 	"github.com/zeromicro/go-zero/rest"
// )

// type contextKey string

// const (
// 	// UserIDKey 用户ID的context key
// 	UserIDKey contextKey = "user_id"
// )

// // AuthMiddleware 认证中间件
// type AuthMiddleware struct {
// 	svcCtx *svc.ServiceContext
// }

// // NewAuthMiddleware 创建认证中间件
// func NewAuthMiddleware(svcCtx *svc.ServiceContext) *AuthMiddleware {
// 	return &AuthMiddleware{
// 		svcCtx: svcCtx,
// 	}
// }

// // Handle 处理请求
// func (m *AuthMiddleware) Handle(r *http.Request) rest.Middleware {
	
// 	// return func(w http.ResponseWriter, r *http.Request) {
// 	// 	// 获取token
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			return rest.Authorize()
			
// 		}

// 		return rest.Authorize()

// 	// 	// 解析token
// 	// 	parts := strings.SplitN(authHeader, " ", 2)
// 	// 	if !(len(parts) == 2 && parts[0] == "Bearer") {
// 	// 		httpx.WriteJson(w, http.StatusUnauthorized, errorx.ErrUnauthorized)
// 	// 		return
// 	// 	}

// 	// 	// 验证token
// 	// 	claims, err := m.svcCtx.Auth.ParseToken(parts[1])
// 	// 	if err != nil {
// 	// 		httpx.WriteJson(w, http.StatusUnauthorized, errorx.ErrUnauthorized)
// 	// 		return
// 	// 	}

// 	// 	// 将用户ID存入context
// 	// 	ctx := context.WithValue(r.Context(), UserIDKey, claims.Id)
// 	// 	next.ServeHTTP(w, r.WithContext(ctx))
// 	// }
// }

// // GetUserID 从context中获取用户ID
// // func GetUserID(ctx context.Context) (uint64, error) {
// // 	userID, ok := ctx.Value(UserIDKey).(uint64)
// // 	if !ok {
// // 		return 0, errorx.ErrUnauthorized
// // 	}
// // 	return userID, nil
// // }

// // Claims 自定义JWT声明
// // type Claims struct {
// // 	UserId int64 `json:"userId"`
// // 	// Username string `json:"username"`
// // 	// Role     string `json:"role"`
// // 	jwt.RegisteredClaims
// // }

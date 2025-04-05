package middleware

// "github.com/golang-jwt/jwt/v4"
// "admin-server/internal/errorx"

// JwtAuthMiddleware JWT认证中间件
// func JwtAuthMiddleware(secret string) rest.Middleware {
// 	return func(next http.HandlerFunc) http.HandlerFunc {
// 		return func(w http.ResponseWriter, r *http.Request) {
// 			// 获取 Authorization header
// 			authHeader := r.Header.Get("Authorization")
// 			if authHeader == "" {
// 				httpx.Error(w, errorx.ErrUnauthorized)
// 				return
// 			}

// 			// 检查 Bearer token
// 			parts := strings.SplitN(authHeader, " ", 2)
// 			if !(len(parts) == 2 && parts[0] == "Bearer") {
// 				httpx.Error(w, errorx.ErrTokenInvalid)
// 				return
// 			}

// 			// 解析 token
// 			claims, err := ParseToken(parts[1], secret)
// 			if err != nil {
// 				if err == jwt.ErrTokenExpired {
// 					httpx.Error(w, errorx.ErrTokenExpired)
// 				} else {
// 					httpx.Error(w, errorx.ErrTokenInvalid)
// 				}
// 				return
// 			}

// 			// 将用户信息存储到上下文中
// 			ctx := r.Context()
// 			ctx = context.WithValue(ctx, "userId", claims.UserId)
// 			// ctx = context.WithValue(ctx, "username", claims.Username)
// 			// ctx = context.WithValue(ctx, "role", claims.Role)
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		}
// 	}
// }

// // GenerateToken 生成JWT token
// func GenerateToken(userId int64, username, role, secret string, expire int64) (string, error) {
// 	claims := Claims{
// 		UserId:   userId,
// 		Username: username,
// 		Role:     role,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Second)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			NotBefore: jwt.NewNumericDate(time.Now()),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(secret))
// }

// // ParseToken 解析JWT token
// func ParseToken(tokenString, secret string) (*Claims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(secret), nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
// 		return claims, nil
// 	}

// 	return nil, errorx.ErrTokenInvalid
// }

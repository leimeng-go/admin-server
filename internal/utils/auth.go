package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func BuildToken(secretKey string, payloads map[string]any, seconds int64) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + seconds
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
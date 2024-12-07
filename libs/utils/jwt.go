package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// secretKey 用于加密 JWT
var secretKey = []byte("secret_key")

// Claims JWT 的 Claims 结构体
type Claims struct {
	Username string `json:"username"`
	UserID   int32  `json:"user_id"`
	jwt.StandardClaims
}

// GenerateJWT 生成 JWT
func GenerateJWT(username string, userID int32) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 过期时间：24小时
	claims := &Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "GoPolyglot",
		},
	}

	// 创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用 secretKey 签名 JWT
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

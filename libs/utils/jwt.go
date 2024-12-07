package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	JWTSecretKey       = "secret-key" // 这里直接写常量，通常会从环境变量或配置文件读取
	JWTExpirationHours = 24
)

// Claims JWT的Claims结构体
type Claims struct {
	Username string `json:"username"`
	UserID   int32  `json:"user_id"`
	jwt.StandardClaims
}

// GenerateJWT 生成JWT Token
func GenerateJWT(username string, userID int32) (string, error) {
	// 设置过期时间
	expirationTime := time.Now().Add(time.Hour * time.Duration(JWTExpirationHours))

	// 创建声明（Claims）
	claims := &Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 设置过期时间
			Issuer:    "GoPolyglot",          // 设置签发者
		},
	}

	// 使用HS256加密算法生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用JWT密钥对Token进行签名
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken 验证JWT Token并返回解析后的用户信息
func ValidateToken(tokenString string) (*Claims, error) {
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法一致
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 类型断言，获取Token中的claims部分
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

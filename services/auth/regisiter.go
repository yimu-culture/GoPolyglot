package auth

import (
	"GoPolyglot/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// VerifyPassword 验证密码是否匹配
func VerifyPassword(storedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword))
}

// HashPassword 对密码进行加密处理
func HashPassword(password string) (string, error) {
	// 使用 bcrypt 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RegisterUser(ctx *gin.Context, username, password string) (*models.User, error) {
	// 首先检查用户名是否已存在
	existingUser, err := models.GetUserByUsername(ctx, username)
	if err == nil && existingUser.ID > 0 {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Username: username,
		Password: hashedPassword,
	}
	user, err := models.CreateUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	// 返回创建的用户
	return user, nil
}

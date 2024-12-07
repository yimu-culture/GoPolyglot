package auth

import (
	"GoPolyglot/models/mysqlDao"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 对密码进行加密处理
func HashPassword(password string) (string, error) {
	// 使用 bcrypt 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RegisterUser(ctx *gin.Context, username, password string) (*mysqlDao.User, error) {
	// 首先检查用户名是否已存在
	existingUser, err := mysqlDao.GetUserByUsername(ctx, username)
	if err == nil && existingUser.ID > 0 {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := mysqlDao.User{
		Username: username,
		Password: hashedPassword,
	}
	user, err := mysqlDao.CreateUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	// 返回创建的用户
	return user, nil
}

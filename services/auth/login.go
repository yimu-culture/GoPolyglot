package auth

import (
	"GoPolyglot/libs/utils"
	"GoPolyglot/models/mysqlDao"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(c *gin.Context, username, password string) (string, error) {
	// 根据用户名查询用户
	user, err := mysqlDao.GetUserByUsername(c, username)
	if err != nil {
		return "", errors.New("user not found")
	}

	// 比对密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	// 密码正确，生成 JWT
	token, err := utils.GenerateJWT(user.Username, user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

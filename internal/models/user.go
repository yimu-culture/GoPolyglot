package models

import (
	"GoPolyglot/libs/dbs" // 引入 dbs 包
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	ID       int32  `gorm:"column:id;primaryKey" json:"id"`
	Username string `gorm:"column:username;unique" json:"username"`
	Password string `gorm:"column:password" json:"-"`
}

// TableName 返回 User 表的名称
func (User) TableName() string {
	return "users"
}

func GetUserByUsername(ctx *gin.Context, username string) (*User, error) {
	var user User
	db := dbs.GMysql["ReelCity"].WithContext(ctx)
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// CreateUser 创建新用户
func CreateUser(ctx *gin.Context, user *User) (*User, error) {
	db := dbs.GMysql["ReelCity"].WithContext(ctx)

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// HashPassword 加密密码
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err // 密码加密失败
	}
	return string(hashedPassword), nil
}

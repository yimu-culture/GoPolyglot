package mysqlDao

import (
	"GoPolyglot/libs/dbs" // 引入 dbs 包
	"context"
	"errors"
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

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	// 从 ReelCity 数据库中查询用户
	db := dbs.GMysql["ReelCity"].WithContext(ctx)
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found") // 用户未找到
	}
	return &user, nil
}

// CreateUser 创建新用户
func CreateUser(ctx context.Context, username, password string) (*User, error) {
	db := dbs.GMysql["ReelCity"].WithContext(ctx)

	// 检查用户名是否已存在
	existingUser, err := GetUserByUsername(ctx, username)
	if err == nil && existingUser.ID > 0 {
		return nil, errors.New("username already exists") // 用户名已存在
	}

	// 加密密码
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err // 密码加密失败
	}

	// 创建新用户
	user := &User{
		Username: username,
		Password: hashedPassword,
	}

	// 保存到数据库
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

// VerifyPassword 验证密码是否匹配
func VerifyPassword(storedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword))
}

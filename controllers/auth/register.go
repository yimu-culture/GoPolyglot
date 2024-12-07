package auth

import (
	"GoPolyglot/libs/common/error_wrapper"
	"GoPolyglot/models/request/auth"
	s_auth "GoPolyglot/services/auth"
	"github.com/gin-gonic/gin"
)

// Register 用户注册接口
func RegisterUser(c *gin.Context) error {
	var request auth.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return error_wrapper.WitheError("Invalid request body", err)
	}

	// 调用 Service 层进行用户注册
	user, err := s_auth.RegisterUser(c, request.Username, request.Password)
	if err != nil {
		return error_wrapper.WitheError(err.Error())
	}

	return error_wrapper.WithSuccess(c, user)
}

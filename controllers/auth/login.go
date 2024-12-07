package auth

import (
	"GoPolyglot/libs/common/error_wrapper"
	"GoPolyglot/models/request/auth"
	res_auth "GoPolyglot/models/response/auth"
	s_auth "GoPolyglot/services/auth"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) error {
	var loginRequest auth.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		return error_wrapper.WitheError(err.Error())
	}

	// 调用 service 层进行登录处理
	token, err := s_auth.LoginService(c, loginRequest.Username, loginRequest.Password)
	if err != nil {
		return error_wrapper.WitheError(err.Error())
	}

	// 返回生成的 JWT
	aa := res_auth.AuthResponse{
		token,
	}
	return error_wrapper.WithSuccessObj(c, aa)
}

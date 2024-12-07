package auth

import (
	"GoPolyglot/libs/common/error_wrapper"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) error {
	return error_wrapper.WithSuccess(c)
}

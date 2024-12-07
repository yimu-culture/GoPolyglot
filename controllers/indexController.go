package controllers

import (
	"GoPolyglot/libs/common/error_wrapper"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) error {
	return error_wrapper.WithSuccess(c)
}

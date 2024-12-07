package error_wrapper

import (
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(ctx *gin.Context) error

func WrapperError(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {

		var err = handler(c)

		if err != nil {
			var errException *ErrorException
			if h, ok := err.(*ErrorException); ok {
				errException = h
			} else if e, ok := err.(error); ok {
				if gin.Mode() == "debug" {
					errException = UnknownError(e.Error())
				} else {
					errException = UnknownError(e.Error())
				}
			} else {
				errException = ServerError()
			}
			errException.Request = c.Request.Method + " " + c.Request.URL.String()
			c.JSON(errException.HttpCode, errException)

			return
		}
	}
}

func WrapperErrors(handler ...HandlerFunc) []gin.HandlerFunc {
	var hf []gin.HandlerFunc

	for _, handlerFunc := range handler {
		tmp := WrapperError(handlerFunc)
		hf = append(hf, tmp)
	}

	return hf

}

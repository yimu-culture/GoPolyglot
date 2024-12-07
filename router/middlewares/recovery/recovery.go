package recovery

import (
	"GoPolyglot/libs/common/error_wrapper"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/outreach-golang/logger"
	"runtime"
)

func panicTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}

// SetUp 捕获请求异常, 并打印日志
func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				s, _ := json.Marshal(panicTrace(r))
				logger.WithContext(c).Error(fmt.Sprintf("system busy: %s", string(s)))
				error_wrapper.WrapperError(func(ctx *gin.Context) error {
					return error_wrapper.WitheError("权益系统繁忙")
				})(c)
			}
		}()
		c.Next()
	}
}

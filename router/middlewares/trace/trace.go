package trace

import (
	"GoPolyglot/libs/configs"
	"bytes"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/outreach-golang/logger"
	"go.uber.org/zap"
	"io/ioutil"
)

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("Traceid")
		if requestId == "" {
			requestId = logger.NewTraceID()
		}
		logger.NewContext(c, zap.String("env", configs.GConfig.Server.Env))
		logger.NewContext(c, zap.String("traceid", requestId))
		logger.NewContext(c, zap.String("local.ip", configs.GConfig.Server.AppName+":"+configs.GConfig.Server.Port))
		logger.NewContext(c, zap.String("request.method", c.Request.Method))
		logger.NewContext(c, zap.String("request.url", c.Request.Host+c.Request.URL.String()))
		logger.NewContext(c, zap.String("request.client.ip", c.ClientIP()))

		headers, _ := jsoniter.Marshal(c.Request.Header)
		logger.NewContext(c, zap.String("request.headers", string(headers)))

		rawData, _ := c.GetRawData()
		logger.NewContext(c, zap.String("request.params", string(rawData)))

		c.Set("traceid", requestId)
		c.Writer.Header().Set("Traceid", requestId)

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))

		c.Next()
	}
}

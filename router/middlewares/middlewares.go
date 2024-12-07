package middleware

import (
	"GoPolyglot/libs/common/error_wrapper"
	"GoPolyglot/libs/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// GetUserFromContext 获取用户信息，从JWT token中提取
func GetUserFromContext(c *gin.Context) (*utils.Claims, error) {
	// 从Authorization头中获取JWT token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return nil, errors.New("authorization header missing")
	}

	// 去掉 "Bearer " 前缀
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// 验证并解析JWT token
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// StrongAuthMiddleware 强认证中间件，依赖于 JWT token 认证
func StrongAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户信息
		claims, err := GetUserFromContext(c)
		if err != nil || claims.UserID <= 0 {
			// 用户未认证或 token 无效
			error_wrapper.WrapperError(func(ctx *gin.Context) error {
				return error_wrapper.WitheError("user session invalid or token expired")
			})(c)
			c.Abort()
			return
		}

		// 将用户ID和用户名存储在上下文中，后续可以通过 c.Get("userID") 和 c.Get("userName") 获取
		c.Set("userID", claims.UserID)
		c.Set("userName", claims.Username)

		// 继续处理请求
		c.Next()
	}
}

// LogMiddleware 日志中间件
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求开始时间
		startTime := time.Now()

		// 从 gin.Context 中获取用户信息
		userID, exists := c.Get("userID")
		if !exists {
			userID = "unknown" // 如果没有用户信息，则标记为"unknown"
		}

		// 获取请求的路径和方法
		path := c.Request.URL.Path
		method := c.Request.Method

		// 处理请求
		c.Next()

		// 计算请求处理时长
		duration := time.Since(startTime)

		// 打印日志
		// 日志格式: [2024-11-21 12:00:00] UserID: userID - POST /tasks/create - 150ms
		// 使用 time.Now().Format 来确保日志时间格式的一致性
		fmt.Printf("[%s] UserID: %v - %s %s - %v\n", time.Now().Format("2006-01-02 15:04:05"), userID, method, path, duration)
	}
}

// 限流设置
const (
	// 限制每个用户每分钟最多请求 10 次
	limitRequestCount = 10
	// 限制周期：单位是分钟
	limitPeriod = time.Minute
)

// 请求记录结构体
type RateLimit struct {
	Count      int
	LastAccess time.Time
}

// 存储用户的请求记录
var requestStore = make(map[string]*RateLimit)

// RateLimitMiddleware 限流中间件
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 gin.Context 中获取用户信息
		userID, exists := c.Get("userID")
		if !exists {
			// 如果没有用户信息，则返回 401 Unauthorized
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid user information",
			})
			c.Abort() // 终止请求
			return
		}

		// 获取当前时间
		now := time.Now()

		// 将 userID 转换为字符串
		userIDStr := fmt.Sprintf("%v", userID)

		// 获取用户的请求记录
		rateLimit, exists := requestStore[userIDStr]
		if !exists {
			rateLimit = &RateLimit{Count: 0, LastAccess: now}
			requestStore[userIDStr] = rateLimit
		}

		// 检查时间是否超过限流周期
		if now.Sub(rateLimit.LastAccess) > limitPeriod {
			// 如果超过了周期，重置请求次数
			rateLimit.Count = 0
			rateLimit.LastAccess = now
		}

		// 判断请求次数是否超过限制
		if rateLimit.Count >= limitRequestCount {
			// 超过限制，返回 429 错误
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			c.Abort() // 终止请求
			return
		}

		// 增加请求计数
		rateLimit.Count++

		// 继续处理请求
		c.Next()
	}
}

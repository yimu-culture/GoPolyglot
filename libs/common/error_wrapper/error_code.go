package error_wrapper

import (
	"GoPolyglot/libs/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/outreach-golang/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type Code int

const (
	SERVER_ERROR Code = 100000 + iota
	NOT_FOUND
	UNKNOWN_ERROR
	PARAMETER_ERROR
	WHITE_LIST
	LEASE_GET_INFO
	LEASE_FORBID
	LEASE_NOT_AUTH
	LEASE_WHITE
	SERVER_LIMITING_ADD_FAIL
	AGGREGATE_NOT_FOUND
	NOT_FIND_EQUITY_INFO
	ENTITY_CARD_INVALID
	ENTITY_CARD_FROZEN
	ENTITY_CARD_CODE_ERROR
	ENTITY_CARD_BALANCE_Not_AVAILABLE
	EquityCenterErr
)

// 账号相关
const (
	ErrAccountNotFound Code = 1101 + iota
	ErrAccountPassFailed
	ErrAccountExist
	ErrTokenEmpty
	ErrTokenInvalid
	ErrTokenExpired
	ErrVerifyCodeInvalid
	ErrAccountForbidden
	ErrVerifyCodeFrequently
	ErrAccountUnsetPass
	ErrAccountNotLogin
	ErrSystemUnknown
	ErrSystemForbidden
	ErrAccountNeedAdd
	ErrAccountAuthErr
	ErrMerchantErr
)

// TODO 业务相关错误

// 其他错误
const (
	ErrOssUpload Code = 1901 + iota
)

var (
	errno         = [...]string{"0", "1"}
	CustomizeCode = map[Code]string{
		SERVER_ERROR:             "系统错误",
		NOT_FOUND:                "404未找到",
		UNKNOWN_ERROR:            "未知",
		PARAMETER_ERROR:          "参数错误",
		WHITE_LIST:               "白名单服务错误",
		LEASE_GET_INFO:           "租户信息错误",
		LEASE_NOT_AUTH:           "无访问权限",
		LEASE_FORBID:             "租户被禁用",
		LEASE_WHITE:              "租户不在白名单中",
		SERVER_LIMITING_ADD_FAIL: "服务限流规则添加失败",
		AGGREGATE_NOT_FOUND:      "服务未找到",
		EquityCenterErr:          "权益中心错误",

		// 账号相关
		ErrAccountNotFound:      "账号未找到",
		ErrAccountPassFailed:    "账号密码错误",
		ErrAccountExist:         "账号已存在",
		ErrTokenEmpty:           "token为空",
		ErrTokenInvalid:         "token无效",
		ErrTokenExpired:         "token已过期",
		ErrVerifyCodeInvalid:    "验证码错误或过期",
		ErrAccountForbidden:     "账号已被禁用",
		ErrVerifyCodeFrequently: "操作频繁，请30分钟后再试",
		ErrAccountUnsetPass:     "账号未设置密码",
		ErrAccountNotLogin:      "用户未登录",
		ErrSystemUnknown:        "未知的系统",
		ErrSystemForbidden:      "系统已被禁用",
		ErrAccountNeedAdd:       "您的账号暂无权限，请联系管理员添加权限",
		ErrAccountAuthErr:       "账号无操作权限",
		ErrMerchantErr:          "供应商不存在或被禁用",

		// 通用
		ErrOssUpload: "文件上传失败",
	}
)

func ErrCodeToStr(code Code) string {
	return common.IntToStr(int(code))
}

// 500 错误处理
func ServerError() *ErrorException {
	return NewErrorException(http.StatusInternalServerError, ErrCodeToStr(SERVER_ERROR), "OBJECT", http.StatusText(http.StatusInternalServerError), "")
}

// 404 错误
func NotFound() *ErrorException {
	return NewErrorException(http.StatusBadRequest, ErrCodeToStr(NOT_FOUND), "OBJECT", http.StatusText(http.StatusNotFound), "")
}

// 未知错误
func UnknownError(message string) *ErrorException {
	return NewErrorException(http.StatusForbidden, ErrCodeToStr(UNKNOWN_ERROR), "OBJECT", message, "")
}

// ParameterError 参数错误
func ParameterError(c *gin.Context, message string) *ErrorException {
	logger.WithContext(c).Error(fmt.Sprintf("WithParameterError"),
		zap.String("message", message),
		zap.String("reqTS", strconv.FormatInt(time.Now().UnixNano()/1e6, 10)))
	return NewErrorException(http.StatusOK, ErrCodeToStr(PARAMETER_ERROR), "OBJECT", message, "")
}

// WithSuccessObj 成功时返回对象
func WithSuccessObj(c *gin.Context, data interface{}) *ErrorException {
	logger.WithContext(c).Info("WithSuccessObj",
		zap.String("data", fmt.Sprintf("%+v", data)),
		zap.String("errcode", errno[0]),
		zap.String("reqTS", strconv.FormatInt(time.Now().UnixNano()/1e6, 10)))
	return NewErrorException(http.StatusOK, errno[0], "OBJECT", "", data)
}

// WithSuccess 成功时返回数组
func WithSuccess(c *gin.Context, data ...interface{}) *ErrorException {
	logger.WithContext(c).Info("WithSuccess",
		zap.String("data", fmt.Sprintf("%+v", data)),
		zap.String("errcode", errno[0]),
		zap.String("reqTS", strconv.FormatInt(time.Now().UnixNano()/1e6, 10)))
	return NewErrorException(http.StatusOK, errno[0], "OBJECT", "", data)
}

// WithErrorObj 错误时返回
func WithErrorObj(c *gin.Context, message string, data interface{}) *ErrorException {
	logger.WithContext(c).Error(fmt.Sprintf("[AIClaimServer] WitheErrorObj error:%s", message),
		zap.String("data", fmt.Sprintf("%+v", data)),
		zap.String("error", message),
		zap.String("reqTS", strconv.FormatInt(time.Now().UnixNano()/1e6, 10)))
	return NewErrorException(http.StatusOK, errno[1], "OBJECT", message, data)
}

// 错误时返回数组
func WitheError(message string, data ...interface{}) *ErrorException {
	return NewErrorException(http.StatusOK, errno[1], "OBJECT", message, data)
}

// 白名单没设置
func WhiteListError(message string, data ...interface{}) *ErrorException {
	return NewErrorException(http.StatusOK, ErrCodeToStr(WHITE_LIST), "OBJECT", message, data)
}

// ErrorCodeObj 自定义错误码
func ErrorCodeObj(c *gin.Context, errorCode Code, data interface{}) *ErrorException {
	logger.WithContext(c).Error(fmt.Sprintf("ErrorCodeObj errcode:%v,errmsg:%s", errorCode, CustomizeCode[errorCode]),
		zap.String("data", fmt.Sprintf("%+v", data)),
		zap.String("errmsg", CustomizeCode[errorCode]),
		zap.String("reqTS", strconv.FormatInt(time.Now().UnixNano()/1e6, 10)))
	return NewErrorException(http.StatusOK, ErrCodeToStr(errorCode), "OBJECT", CustomizeCode[errorCode], data)
}

// ErrorCodeMsgObj 自定义错误码 自定义错误信息
func ErrorCodeMsgObj(c *gin.Context, errorCode Code, errMsg string, data interface{}) *ErrorException {
	if errMsg == "" {
		errMsg = CustomizeCode[errorCode]
	}
	logger.WithContext(c).Error(fmt.Sprintf("ErrorCodeObj errcode:%v,errmsg:%s", errorCode, CustomizeCode[errorCode]),
		zap.String("data", fmt.Sprintf("%+v", data)),
		zap.String("errmsg", errMsg),
		zap.String("reqTS", strconv.FormatInt(time.Now().UnixNano()/1e6, 10)))
	return NewErrorException(http.StatusOK, ErrCodeToStr(errorCode), "OBJECT", CustomizeCode[errorCode], data)
}

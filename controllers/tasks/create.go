package tasks

import (
	"GoPolyglot/libs/common/error_wrapper"
	task_request "GoPolyglot/models/request/task"
	task_service "GoPolyglot/services/task"
	"github.com/gin-gonic/gin"
)

// CreateTranslationTask 处理翻译任务创建的请求
func CreateTask(c *gin.Context) error {
	var createRequest task_request.CreateTranslationTaskRequest

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		return error_wrapper.WitheError(err.Error())
	}

	// 从上下文中获取当前用户的信息
	userID, userExists := c.Get("userID")
	if !userExists {
		return error_wrapper.WitheError("User session invalid")
	}

	// 类型断言，将 userID 转换为 int32
	userIDInt, valid := userID.(int32)
	if !valid {
		return error_wrapper.WitheError("Invalid user ID format")
	}

	// 调用服务层的创建任务方法
	translationTask, err := task_service.CreateTask(c, userIDInt, createRequest.SourceLang, createRequest.TargetLang, createRequest.SourceDoc)
	if err != nil {
		return error_wrapper.WitheError("Failed to create translation task")
	}

	// 返回成功的响应，包含任务信息
	return error_wrapper.WithSuccess(c, translationTask)
}

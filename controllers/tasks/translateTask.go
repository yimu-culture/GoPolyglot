package tasks

import (
	"GoPolyglot/libs/common/error_wrapper"
	"GoPolyglot/services/task"
	"github.com/gin-gonic/gin"
	"strconv"
)

// TranslateTask 处理翻译任务的执行
func TranslateTask(ctx *gin.Context) error {
	// 从 URL 中获取任务ID
	taskIDStr := ctx.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return error_wrapper.WitheError("User session invalid1")
	}

	// 调用 Service 层来启动翻译任务
	task, err := task.StartTranslationTask(ctx, int32(taskID))
	if err != nil {
		return error_wrapper.WitheError("User session invalid2")
	}

	// 返回任务的状态及相关信息
	return error_wrapper.WithSuccessObj(ctx, map[string]interface{}{
		"message": "Translation task started successfully",
		"task":    task,
	})
}

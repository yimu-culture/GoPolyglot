package tasks

import (
	"GoPolyglot/global" // 引入全局包
	"GoPolyglot/libs/common/error_wrapper"
	"github.com/gin-gonic/gin"
	"strconv"
)

// TranslateTask 处理翻译任务
func TranslateTask(ctx *gin.Context) error {
	// 从 URL 中获取任务ID
	taskIDStr := ctx.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return error_wrapper.WitheError("Invalid task ID")
	}

	// 提交任务到全局协程池
	global.SubmitTask(int32(taskID))

	// 返回成功消息
	return error_wrapper.WithSuccessObj(ctx, map[string]interface{}{
		"message": "Translation task started",
		"task_id": taskID,
	})
}

package tasks

import (
	"GoPolyglot/libs/common/error_wrapper"
	"GoPolyglot/services/task"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetTaskStatus(ctx *gin.Context) error {
	// 从 URL 中获取任务ID
	taskIDStr := ctx.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return error_wrapper.WitheError("Invalid task ID")
	}

	// 获取任务状态
	status, err := task.GetTaskStatus(int32(taskID))
	if err != nil {
		return error_wrapper.WitheError("Failed to retrieve task status")
	}

	// 返回任务状态
	return error_wrapper.WithSuccessObj(ctx, map[string]interface{}{
		"task_id": taskID,
		"status":  status,
	})
}

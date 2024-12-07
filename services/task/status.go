package task

import (
	"GoPolyglot/models/mysqlDao"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetTaskStatus 获取任务的状态
func GetTaskStatus(ctx *gin.Context, taskID int32) (string, error) {
	task, err := mysqlDao.GetTranslationTaskByID(ctx, taskID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve task: %v", err)
	}
	return task.Status, nil
}

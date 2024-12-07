package task

import (
	"GoPolyglot/models/mysqlDao"
	"fmt"
)

// GetTaskStatus 获取任务的状态
func GetTaskStatus(taskID int32) (string, error) {
	task, err := mysqlDao.GetTranslationTaskByID(nil, taskID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve task: %v", err)
	}
	return task.Status, nil
}

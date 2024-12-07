package tasks

import (
	"GoPolyglot/libs/common/error_wrapper"
	"GoPolyglot/models/mysqlDao"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
)

func DownloadTranslation(ctx *gin.Context) error {
	// 从 URL 中获取任务ID
	taskIDStr := ctx.Param("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return error_wrapper.WitheError("Invalid task ID")
	}

	// 获取翻译内容
	task, err := mysqlDao.GetTranslationTaskByID(ctx, int32(taskID))
	if err != nil {
		return error_wrapper.WitheError("Failed to retrieve translated document")
	}

	// 假设 translatedDoc 是文件的路径，返回文件下载
	filePath := task.TranslatedDoc // 文件路径，例如 "/tmp/translated_task_8.txt"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return error_wrapper.WitheError("File not found")
	}

	// 设置响应头以便浏览器下载文件
	ctx.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	ctx.Header("Content-Type", "application/octet-stream")

	// 返回文件
	//if err := ctx.File(filePath); err != nil {
	//	return error_wrapper.WrapperError("Failed to send file")
	//}
	return error_wrapper.WithSuccessObj(ctx, ctx)
}

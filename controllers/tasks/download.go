package tasks

import (
	"GoPolyglot/libs/common/error_wrapper"
	"GoPolyglot/models/mysqlDao"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

	// 假设任务中保存的翻译内容
	translatedDoc := task.TranslatedDoc

	// 如果翻译内容为空，返回错误
	if translatedDoc == "" {
		return error_wrapper.WitheError("No translated content found")
	}

	// 创建临时文件来存储翻译结果
	tempDir := "./tmp"
	// 创建目录（如果没有的话）
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return error_wrapper.WitheError("Failed to create temporary directory")
	}

	// 文件名带上时间戳，避免文件重名
	fileName := "translated_task_" + strconv.Itoa(taskID) + "_" + time.Now().Format("20060102150405") + ".txt"
	filePath := filepath.Join(tempDir, fileName)

	// 使用 os.WriteFile 将翻译内容写入文件
	err = os.WriteFile(filePath, []byte(translatedDoc), 0644)
	if err != nil {
		return error_wrapper.WitheError("Failed to write translated content to file")
	}

	// 设置响应头以便浏览器下载文件
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/octet-stream")

	// 返回文件
	ctx.File(filePath) // 这里直接调用，不会返回错误

	// 发送成功响应并删除文件
	defer os.Remove(filePath) // 清理临时文件

	return nil
}

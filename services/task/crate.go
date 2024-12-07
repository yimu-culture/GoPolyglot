package task

import (
	"GoPolyglot/models/mysqlDao"
	"github.com/gin-gonic/gin"
)

// CreateTask 创建翻译任务
func CreateTask(c *gin.Context, userID int32, sourceLang, targetLang, sourceDoc string) (*mysqlDao.TranslationTask, error) {
	// 创建翻译任务对象
	task := &mysqlDao.TranslationTask{
		UserID:     userID,
		SourceLang: sourceLang,
		TargetLang: targetLang,
		Status:     "pending", // 初始状态是待翻译
		SourceDoc:  sourceDoc,
	}

	// 调用 model 层创建任务
	return mysqlDao.CreateTask(c, task)
}

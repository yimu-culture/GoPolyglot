package mysqlDao

import (
	"GoPolyglot/libs/dbs"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TranslationTask 翻译任务模型
type TranslationTask struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 任务的唯一标识符，主键，自增
	UserID        int32     `gorm:"column:user_id" json:"user_id"`                // 发起翻译任务的用户ID
	SourceLang    string    `gorm:"column:source_lang" json:"source_lang"`        // 源语言标识符
	TargetLang    string    `gorm:"column:target_lang" json:"target_lang"`        // 目标语言标识符
	Status        string    `gorm:"column:status" json:"status"`                  // 当前任务状态：待翻译、翻译中、已完成
	SourceDoc     string    `gorm:"column:source_doc" json:"source_doc"`          // 源文档路径或内容
	TranslatedDoc string    `gorm:"column:translated_doc" json:"translated_doc"`  // 翻译后文档的路径或内容
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`          // 任务创建时间
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`          // 任务最后更新时间
}

// TableName 重写表名
func (TranslationTask) TableName() string {
	return "translation_tasks"
}

// CreateTranslationTask 创建翻译任务
func CreateTranslationTask(ctx *gin.Context, task *TranslationTask) (*TranslationTask, error) {
	result := dbs.GMysql["ReelCity"].Create(task)
	if result.Error != nil {
		return nil, result.Error // 返回错误信息
	}
	return task, nil
}

// GetTranslationTaskByID 根据任务ID查询翻译任务
func GetTranslationTaskByID(ctx *gin.Context, taskID int32) (*TranslationTask, error) {
	var task TranslationTask
	result := dbs.GMysql["ReelCity"].First(&task, taskID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("task ID %d not found", taskID) // 如果任务未找到，返回错误
		}
		return nil, result.Error // 其他错误，直接返回
	}
	return &task, nil
}

// UpdateTranslationTask 更新翻译任务
func UpdateTranslationTask(ctx *gin.Context, taskID int32, updateData map[string]interface{}) (*TranslationTask, error) {
	var task TranslationTask
	result := dbs.GMysql["ReelCity"].First(&task, taskID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("task ID %d not found", taskID)
		}
		return nil, result.Error
	}

	// 更新任务字段
	result = dbs.GMysql["ReelCity"].Model(&task).Updates(updateData)
	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

// GetAllTranslationTasksByUser 获取用户所有翻译任务
func GetAllTranslationTasksByUser(ctx *gin.Context, userID int32) ([]TranslationTask, error) {
	var tasks []TranslationTask
	result := dbs.GMysql["ReelCity"].Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

// DeleteTranslationTask 根据任务ID删除翻译任务
func DeleteTranslationTask(ctx *gin.Context, taskID int32) error {
	var task TranslationTask
	result := dbs.GMysql["ReelCity"].First(&task, taskID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("task ID %d not found", taskID)
		}
		return result.Error
	}

	// 删除任务记录
	result = dbs.GMysql["ReelCity"].Delete(&task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

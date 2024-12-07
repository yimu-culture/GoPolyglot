package mysqlDao

import (
	"GoPolyglot/libs/dbs"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// TranslationTask 翻译任务模型
type TranslationTask struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID        int32     `gorm:"column:user_id" json:"user_id"`
	SourceLang    string    `gorm:"column:source_lang" json:"source_lang"`
	TargetLang    string    `gorm:"column:target_lang" json:"target_lang"`
	Status        string    `gorm:"column:status" json:"status"`
	SourceDoc     string    `gorm:"column:source_doc" json:"source_doc"`
	TranslatedDoc string    `gorm:"column:translated_doc" json:"translated_doc"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName 重写表名
func (TranslationTask) TableName() string {
	return "translation_tasks"
}

func CreateTask(ctx *gin.Context, task *TranslationTask) (*TranslationTask, error) {
	db := dbs.GMysql["ReelCity"].WithContext(ctx)

	if err := db.Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// GetTranslationTaskByID 根据任务ID查询任务
func GetTranslationTaskByID(ctx *gin.Context, taskID int32) (*TranslationTask, error) {
	var task TranslationTask
	result := dbs.GMysql["ReelCity"].WithContext(ctx).First(&task, taskID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("task ID %d not found", taskID)
		}
		return nil, result.Error
	}
	return &task, nil
}

// UpdateTranslationTask 更新翻译任务状态
func UpdateTranslationTask(ctx *gin.Context, taskID int32, updateData map[string]interface{}) (*TranslationTask, error) {
	var task TranslationTask
	result := dbs.GMysql["ReelCity"].WithContext(ctx).First(&task, taskID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("task ID %d not found", taskID)
		}
		return nil, result.Error
	}

	result = dbs.GMysql["ReelCity"].Model(&task).Updates(updateData)
	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

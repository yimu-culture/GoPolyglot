package mysqlDao

import (
	"GoPolyglot/libs/dbs"
	"github.com/gin-gonic/gin"
	"time"
)

// TranslationTask 翻译任务模型
type TranslationTask struct {
	ID            int32     `gorm:"column:id;primaryKey" json:"id"`              // 任务的唯一标识符
	UserID        int32     `gorm:"column:user_id" json:"user_id"`               // 发起翻译任务的用户ID
	SourceLang    string    `gorm:"column:source_lang" json:"source_lang"`       // 源语言标识符
	TargetLang    string    `gorm:"column:target_lang" json:"target_lang"`       // 目标语言标识符
	Status        string    `gorm:"column:status" json:"status"`                 // 当前任务状态：待翻译、翻译中、已完成
	SourceDoc     string    `gorm:"column:source_doc" json:"source_doc"`         // 源文档路径或内容
	TranslatedDoc string    `gorm:"column:translated_doc" json:"translated_doc"` // 翻译后文档的路径或内容
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`         // 任务创建时间
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`         // 任务最后更新时间
}

// TableName 重写表名
func (TranslationTask) TableName() string {
	return "translation_tasks"
}

// CreateTranslationTask 创建翻译任务
func CreateTranslationTask(ctx *gin.Context, task *TranslationTask) (*TranslationTask, error) {
	result := dbs.GMysql["ReelCity"].Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

package mysqlDao

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	Status      string    `json:"status"`
	SourceLang  string    `json:"source_lang"`
	TargetLang  string    `json:"target_lang"`
	Content     string    `json:"content"`
	Result      string    `json:"result"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

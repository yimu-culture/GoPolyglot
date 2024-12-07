package task

type CreateTranslationTaskRequest struct {
	SourceLang string `json:"source_lang" binding:"required"`
	TargetLang string `json:"target_lang" binding:"required"`
	// 这可以是文件存储系统中的文件路径或URL，也可以是文档的直接内容
	SourceDoc string `json:"source_doc" binding:"required"`
}

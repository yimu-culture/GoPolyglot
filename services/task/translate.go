package task

import (
	"GoPolyglot/models/mysqlDao"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// StartTranslationTask 启动翻译任务
// 该方法负责启动翻译任务并调用外部LLM API来执行翻译
func StartTranslationTask(ctx *gin.Context, taskID int32) (*mysqlDao.TranslationTask, error) {
	// 1. 获取翻译任务
	task, err := mysqlDao.GetTranslationTaskByID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve task: %v", err)
	}

	// 2. 更新任务状态为 "翻译中"
	_, err = mysqlDao.UpdateTranslationTask(ctx, taskID, map[string]interface{}{
		"status": "in_progress",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update task status: %v", err)
	}

	// 3. 调用外部LLM翻译API
	translatedDoc, err := callTranslationAPI(task.SourceDoc, task.SourceLang, task.TargetLang)
	if err != nil {
		// 如果翻译失败，将任务状态更新为 "翻译失败"
		_, errUpdate := mysqlDao.UpdateTranslationTask(ctx, taskID, map[string]interface{}{
			"status": "fail",
		})
		if errUpdate != nil {
			log.Printf("failed to update task status to 'failed': %v", errUpdate)
		}
		return nil, fmt.Errorf("translation failed: %v", err)
	}

	// 4. 更新任务状态为 "已完成" 并保存翻译后的文档路径或内容
	_, err = mysqlDao.UpdateTranslationTask(ctx, taskID, map[string]interface{}{
		"status":         "completed",
		"translated_doc": translatedDoc,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update task with translated document: %v", err)
	}

	// 5. 返回已更新的任务
	return task, nil
}

// callTranslationAPI 调用外部翻译API进行翻译
// 该方法将文档内容、源语言和目标语言传递给外部翻译服务
func callTranslationAPI(sourceDoc, sourceLang, targetLang string) (string, error) {
	// 这里假设我们使用一个外部API进行翻译，具体实现可根据实际API调整
	// 这里使用一个假的API URL
	apiURL := "https://example.com/translate"

	// 构造请求的 payload
	requestPayload := map[string]interface{}{
		"source_doc":  sourceDoc,
		"source_lang": sourceLang,
		"target_lang": targetLang,
	}
	payload, err := json.Marshal(requestPayload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request payload: %v", err)
	}

	// 向外部API发送POST请求
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("failed to call translation API: %v", err)
	}
	defer resp.Body.Close()

	// 检查HTTP响应
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("translation API returned non-OK status: %s", resp.Status)
	}

	// 解析翻译结果
	var apiResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return "", fmt.Errorf("failed to decode API response: %v", err)
	}

	// 返回翻译后的文档内容
	translatedDoc, ok := apiResponse["translated_doc"].(string)
	if !ok {
		return "", fmt.Errorf("translation API response does not contain 'translated_doc'")
	}

	return translatedDoc, nil
}

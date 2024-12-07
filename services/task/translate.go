package task

import (
	"GoPolyglot/models/mysqlDao"
	"fmt"
	"log"
	"sync"
)

// WorkerPool 协程池
type WorkerPool struct {
	taskQueue  chan int32
	numWorkers int
	wg         sync.WaitGroup
}

// Start 启动工作池
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		go wp.worker(i)
	}
}

// worker 处理翻译任务
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for taskID := range wp.taskQueue {
		err := processTranslationTask(taskID)
		if err != nil {
			log.Printf("Worker %d failed to process task %d: %v", id, taskID, err)
		}
	}
}

// SubmitTask 提交任务
func (wp *WorkerPool) SubmitTask(taskID int32) {
	wp.wg.Add(1)
	wp.taskQueue <- taskID
}

// Wait 等待所有任务完成
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

// processTranslationTask 处理翻译任务
func processTranslationTask(taskID int32) error {
	_, err := mysqlDao.GetTranslationTaskByID(nil, taskID)
	if err != nil {
		return fmt.Errorf("failed to retrieve task: %v", err)
	}

	// 更新任务状态
	_, err = mysqlDao.UpdateTranslationTask(nil, taskID, map[string]interface{}{
		"status": "in_progress",
	})
	if err != nil {
		return fmt.Errorf("failed to update task status: %v", err)
	}

	// 模拟翻译过程
	translatedDoc := "Translated document content" // 假设翻译成功
	_, err = mysqlDao.UpdateTranslationTask(nil, taskID, map[string]interface{}{
		"status":         "completed",
		"translated_doc": translatedDoc,
	})
	if err != nil {
		log.Printf("failed to update task with translated document: %v", err)
	}

	return nil
}

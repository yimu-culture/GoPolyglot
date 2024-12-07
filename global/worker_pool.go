package global

import (
	"GoPolyglot/models/mysqlDao"
	"fmt"
	"log"
	"sync"
)

// WorkerPool 协程池结构
type WorkerPool struct {
	taskQueue  chan int32
	numWorkers int
	wg         sync.WaitGroup
}

// GlobalWorkerPool 用于管理全局单例的 WorkerPool
var GlobalWorkerPool *WorkerPool

// InitWorkerPool 初始化全局协程池
func InitWorkerPool(numWorkers int) {
	GlobalWorkerPool = &WorkerPool{
		taskQueue:  make(chan int32, 100),
		numWorkers: numWorkers,
	}

	for i := 0; i < numWorkers; i++ {
		go GlobalWorkerPool.worker(i)
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

// SubmitTask 提交任务到协程池
func SubmitTask(taskID int32) {
	GlobalWorkerPool.wg.Add(1)
	GlobalWorkerPool.taskQueue <- taskID
}

// Wait 等待所有任务完成
func Wait() {
	GlobalWorkerPool.wg.Wait()
}

// processTranslationTask 模拟翻译任务的处理
func processTranslationTask(taskID int32) error {
	// 从数据库获取任务
	_, err := mysqlDao.GetTranslationTaskByID(nil, taskID)
	if err != nil {
		return fmt.Errorf("failed to retrieve task: %v", err)
	}

	// 更新任务状态为 "in_progress"
	_, err = mysqlDao.UpdateTranslationTask(nil, taskID, map[string]interface{}{
		"status": "in_progress",
	})
	if err != nil {
		return fmt.Errorf("failed to update task status: %v", err)
	}

	// 模拟翻译
	translatedDoc := "Translated document content"
	_, err = mysqlDao.UpdateTranslationTask(nil, taskID, map[string]interface{}{
		"status":         "completed",
		"translated_doc": translatedDoc,
	})
	if err != nil {
		log.Printf("failed to update task with translated document: %v", err)
	}

	return nil
}

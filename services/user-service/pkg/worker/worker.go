package worker

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type WorkerPool struct {
	workerName string
	workers    int
	taskQueue  chan TaskQueue
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewWorkerPool(workerName string, workers int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		workerName: workerName,
		workers:    workers,
		taskQueue:  make(chan TaskQueue),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func NewWorkerPoolTaskQueue(workerName string, workers int, queueSize int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		workerName: workerName,
		workers:    workers,
		taskQueue:  make(chan TaskQueue, queueSize),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (wp *WorkerPool) Start() {
	log.Println(fmt.Sprintf("worker pool starting %s with workers %d", wp.workerName, wp.workers))
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for {
		select {
		case task := <-wp.taskQueue:
			if err := task.Handler(task.Ctx, task.Delivery); err != nil {
				log.Printf("Worker Queue %s : Task Error: %v", wp.workerName, err)
				return
			}
		case <-wp.ctx.Done():
			log.Printf("Worker %d stopped", id)
			return
		}
	}
}

func (wp *WorkerPool) Stop() {
	wp.cancel()
	close(wp.taskQueue)
	wp.wg.Wait()
}

package worker

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type WorkerPool struct {
	workerName        string
	workers           int
	rabbitmqTaskQueue chan RabbitMQTaskQueue
	kafkaTaskQueue    chan KafkaTaskQueue
	wg                sync.WaitGroup
	ctx               context.Context
	cancel            context.CancelFunc
}

func NewWorkerPool(workerName string, workers int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		workerName:        workerName,
		workers:           workers,
		rabbitmqTaskQueue: make(chan RabbitMQTaskQueue),
		ctx:               ctx,
		cancel:            cancel,
	}
}

func NewWorkerPoolRabbitMQTaskQueue(workerName string, workers int, queueSize int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		workerName:        workerName,
		workers:           workers,
		rabbitmqTaskQueue: make(chan RabbitMQTaskQueue, queueSize),
		kafkaTaskQueue:    nil,
		ctx:               ctx,
		cancel:            cancel,
	}
}

func NewWorkerPoolKafkaTaskQueue(workerName string, workers int, queueSize int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		workerName:        workerName,
		workers:           workers,
		rabbitmqTaskQueue: nil,
		kafkaTaskQueue:    make(chan KafkaTaskQueue, queueSize),
		ctx:               ctx,
		cancel:            cancel,
	}
}

func (wp *WorkerPool) AddRabbitMQTaskQueue(task RabbitMQTaskQueue) {
	select {
	case wp.rabbitmqTaskQueue <- task:
		log.Printf("RabbitMQ Queue Name %s : Task Added", task.QueueName)
	default:
		log.Printf("Worker %s : %s  is full : %v", wp.workerName, task.QueueName, task)
	}
}

func (wp *WorkerPool) AddKafkaTaskQueue(task KafkaTaskQueue) {
	select {
	case wp.kafkaTaskQueue <- task:
		log.Printf("Kafka Topic Name : %s Task Added", task.Message)
	default:
		log.Printf("Worker %s : %s  is full : %v", wp.workerName, task.Message, task)
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
		case task := <-wp.rabbitmqTaskQueue:
			if err := task.Handler(task.Ctx, task.Delivery); err != nil {
				log.Printf("RabbitMQ Worker Queue %s : Task Error: %v", wp.workerName, err)
				return
			}
		case task := <-wp.kafkaTaskQueue:
			if err := task.Handler(task.Ctx, task.Message); err != nil {
				log.Printf("kafka Worker Queue %s : Task Error: %v", wp.workerName, err)
				return
			}
		case <-wp.ctx.Done():
			log.Printf("Worker %s :  %d stopped", wp.workerName, id)
			return
		}
	}
}

func (wp *WorkerPool) Stop() {
	wp.cancel()
	if wp.rabbitmqTaskQueue != nil {
		close(wp.rabbitmqTaskQueue)
	}
	if wp.kafkaTaskQueue != nil {
		close(wp.kafkaTaskQueue)
	}
	wp.wg.Wait()
}

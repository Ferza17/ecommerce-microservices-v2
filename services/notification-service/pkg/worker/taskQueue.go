package worker

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type TaskQueue struct {
	QueueName string
	Ctx       context.Context
	Delivery  *amqp091.Delivery
	Handler   func(ctx context.Context, d *amqp091.Delivery) error
}

func (wp *WorkerPool) AddTaskQueue(task TaskQueue) {
	select {
	case wp.taskQueue <- task:
	default:
		log.Printf("Worker %s : %s  is full : %v", wp.workerName, task.QueueName, task)
	}
}

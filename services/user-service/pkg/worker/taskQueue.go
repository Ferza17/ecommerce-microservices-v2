package worker

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMQTaskQueue struct {
	QueueName string
	Ctx       context.Context
	Delivery  *amqp091.Delivery
	Handler   func(ctx context.Context, d *amqp091.Delivery) error
}

func (wp *WorkerPool) AddTaskQueue(task RabbitMQTaskQueue) {
	select {
	case wp.rabbitmqTaskQueue <- task:
		log.Printf("Queue Name %s : Task Added", task.QueueName)
	default:
		log.Printf("Worker %s : %s  is full : %v", wp.workerName, task.QueueName, task)
	}
}

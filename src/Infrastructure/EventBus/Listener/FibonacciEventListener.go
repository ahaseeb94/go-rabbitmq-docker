package Listener

import (
	"go-rabbitmq-docker/src/Application/Service"
	"go-rabbitmq-docker/src/Infrastructure/EventBus"
	"os"
)

// FibonacciEventListener -> FibonacciEventListener
type FibonacciEventListener struct {
	*EventBus.RabbitMQListener //inheritance
}

// NewFibonacciEventListener : NewFibonacciEventListener
func NewFibonacciEventListener() *FibonacciEventListener {
	amqpHost := os.Getenv("AMQP_HOST")
	amqpUser := os.Getenv("AMQP_USER")
	amqpPass := os.Getenv("AMQP_PASS")
	queueName := os.Getenv("AMQP_FIBONACCI_QUEUE")
	return &FibonacciEventListener{
		RabbitMQListener: &EventBus.RabbitMQListener{
			NewService: func() EventBus.EventListener {
				return &Service.FibonacciListenerService{}
			},
			AmqpHost:  amqpHost,
			AmqpUser:  amqpUser,
			AmqpPass:  amqpPass,
			QueueName: queueName,
		},
	}
}

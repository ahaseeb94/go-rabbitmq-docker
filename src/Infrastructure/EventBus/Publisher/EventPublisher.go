package Publisher

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

type EventPublisher struct{}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{}
}

func (s *EventPublisher) Publish(message []byte, exchange string) {
	amqpHost := os.Getenv("AMQP_HOST")
	amqpUser := os.Getenv("AMQP_USER")
	amqpPass := os.Getenv("AMQP_PASS")
	port := 5672
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", amqpUser, amqpPass, amqpHost, port))
	FailOnError(err, "Failed to connect to RabbitMQ")

	rabbitMQPublisher := NewRabbitMQPublisher(conn, exchange)
	rabbitMQPublisher.Publish(message)
}

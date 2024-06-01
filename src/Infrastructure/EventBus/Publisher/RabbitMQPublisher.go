package Publisher

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type RabbitMQPublisher struct {
	conn     *amqp.Connection
	exchange string
}

func NewRabbitMQPublisher(conn *amqp.Connection, exchange string) *RabbitMQPublisher {
	return &RabbitMQPublisher{
		conn:     conn,
		exchange: exchange,
	}
}
func (p *RabbitMQPublisher) Publish(message []byte) {
	defer p.conn.Close()
	ch, err := p.conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	ctx_amqp, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx_amqp,
		p.exchange, // exchange
		"*",        // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	FailOnError(err, "Failed to publish a message")
	fmt.Println(" [x] Sent " + string(message))
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

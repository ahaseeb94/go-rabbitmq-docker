package EventBus

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

// RabbitMQListener -> RabbitMQListener
type RabbitMQListener struct {
	NewService func() EventListener
	AmqpHost   string
	AmqpUser   string
	AmqpPass   string
	QueueName  string
}

// Listen : Listen
func (p *RabbitMQListener) Listen() {
	prefetchCount := 5
	port := 5672
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", p.AmqpUser, p.AmqpPass, p.AmqpHost, port))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	err = channel.Qos(
		prefetchCount, // prefetch count
		0,             // prefetch size (0 means unlimited)
		false,         // global (false means the QoS settings apply to the current channel)
	)
	if err != nil {
		log.Fatalf("Failed to set QoS: %s", err)
	}

	msgs, err := channel.Consume(
		p.QueueName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	for i := 0; i < prefetchCount; i++ {
		go func() {
			for d := range msgs {
				p.Process(d)
			}
		}()
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}

func (p *RabbitMQListener) Process(d amqp.Delivery) {
	log.Printf("Received a message: %s", d.Body)
	service := p.NewService()
	service.Listen(d.Body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

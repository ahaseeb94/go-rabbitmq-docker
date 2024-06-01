package Service

import (
	"encoding/json"
	"fmt"
	"go-rabbitmq-docker/src/Application/Request"
	"go-rabbitmq-docker/src/Infrastructure/EventBus/Publisher"
	"os"
)

type Fibonacci struct {
	Number int `json:"number"`
}

type FibonacciService struct{}

func NewFibonacciService() *FibonacciService {
	return &FibonacciService{}
}

func (s *FibonacciService) Execute(request *Request.FibonacciRequestDraft) {
	exchange := os.Getenv("AMQP_FIBONACCI_EXCHANGE")
	msg := Fibonacci{Number: request.Number}
	msgJson, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Failed to marshal number to JSON.")
	}

	s.createEvent(msgJson, exchange)
}

func (s *FibonacciService) createEvent(eventBody []byte, exchange string) {
	eventPublisher := Publisher.NewEventPublisher()
	eventPublisher.Publish(eventBody, exchange)
}

package Command

import (
	"go-rabbitmq-docker/src/Infrastructure/EventBus/Listener"
)

// FibonacciCommandService -> FibonacciCommandService
type FibonacciCommandService struct {
	listener *Listener.FibonacciEventListener
}

// NewFibonacciCommandService : NewFibonacciCommandService
func NewFibonacciCommandService() *FibonacciCommandService {
	return &FibonacciCommandService{
		listener: Listener.NewFibonacciEventListener(),
	}
}

func (s *FibonacciCommandService) Execute() {
	s.listener.Listen()
}

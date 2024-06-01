package Service

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
)

type FibonacciListenerService struct{}

func NewFibonacciListenerService() *FibonacciListenerService {
	return &FibonacciListenerService{}
}

func (s *FibonacciListenerService) Listen(request []byte) (map[string]interface{}, error) {
	var msg Fibonacci
	err := json.Unmarshal(request, &msg)
	if err != nil {
		return nil, err
	}
	num := msg.Number
	log.Printf("Received number: %d", num)

	result := fmt.Sprintf("The %dth Fibonacci number is: %s", num, s.fibonacci(num))
	log.Printf(result + "\n")

	response := map[string]interface{}{
		"status":  "success",
		"message": result,
	}
	return response, nil
}

// fibonacci calculates the nth Fibonacci number.
func (s *FibonacciListenerService) fibonacci(n int) *big.Int {
	prev, curr := big.NewInt(0), big.NewInt(1)

	for i := 0; i < n; i++ {
		next := new(big.Int).Add(prev, curr)
		prev, curr = curr, next
	}

	return prev
}

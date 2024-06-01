package Listener

import (
	"go-rabbitmq-docker/src/Application/Service"
	"reflect"
	"testing"
)

func TestFibonacciListenerService(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		service := Service.NewFibonacciListenerService()
		request := []byte(`{"number":9}`)
		expected := map[string]interface{}{
			"status":  "success",
			"message": "The 9th Fibonacci number is: 34",
		}

		response, err := service.Listen(request)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(response, expected) {
			t.Errorf("unexpected response: got %v, want %v", response, expected)
		}
	})

}

package Handler

import (
	"github.com/gin-gonic/gin"
	"go-rabbitmq-docker/src/Application/Request"
	"go-rabbitmq-docker/src/Application/Responder/Core"
	"go-rabbitmq-docker/src/Application/Service"
	"net/http"
)

const (
	EventSuccessCode          = "GO-ES-001"
	QueueSuccessMessageFormat = "Event successfully added in the queue!"
)

type FibonacciHandler struct {
	service *Service.FibonacciService
}

func NewFibonacciHandler() *FibonacciHandler {
	return &FibonacciHandler{
		service: Service.NewFibonacciService(),
	}
}

func (h *FibonacciHandler) Execute(c *gin.Context) {
	var eventRequest Request.FibonacciRequestDraft

	if err := c.ShouldBind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.service.Execute(&eventRequest)

	content := map[string]interface{}{
		"message": QueueSuccessMessageFormat,
	}

	response := Core.UnformedResponseDraft(EventSuccessCode, http.StatusOK, content)

	c.JSON(response.HTTPCode, response)
}

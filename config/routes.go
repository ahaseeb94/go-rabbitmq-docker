package config

import (
	"github.com/gin-gonic/gin"
	"go-rabbitmq-docker/src/Application/Handler"
)

type Routes struct{}

func NewRoutes() *Routes {
	return &Routes{}
}

func (r *Routes) SetUpRoutes(app *gin.Engine) {
	apiRoutes := app.Group("api")
	{
		fibonacciHandler := Handler.NewFibonacciHandler()
		apiRoutes.POST("/publish", fibonacciHandler.Execute)
	}
}

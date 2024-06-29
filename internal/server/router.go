package server

import (
	"go-api-event/configs"
	"go-api-event/internal/controllers"
	rabbitmq "go-api-event/internal/rabbitMQ"
)

func (s server) registerRoutes() {
	cfg := configs.GetConfig()

	var rabbitMQService, err = rabbitmq.NewRabbitMQService(s.AmqpConn, cfg.RabbitMQ)
	if err != nil {
		panic(err)
	}

	c := controllers.NewEventController(s.Logger, rabbitMQService)
	s.router.POST("/v1/event/publish", c.Publish)
}

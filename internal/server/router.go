package server

import (
	"go-api-event/configs"
	"go-api-event/internal/controllers"
	"go-api-event/internal/db/redisdb"
	rabbitmq "go-api-event/internal/rabbitMQ"
	"go-api-event/internal/services"
)

func (s server) registerRoutes() {
	cfg := configs.GetConfig()

	var rabbitMQService, err = rabbitmq.NewRabbitMQService(s.AmqpConn, cfg.RabbitMQ)
	if err != nil {
		panic(err)
	}

	var eventDb = redisdb.NewEventRedisDB(s.RedisConn)
	var service = services.NewService(rabbitMQService, eventDb)

	c := controllers.NewEventController(s.Logger, service)
	s.router.POST("/v1/event/publish", c.Publish)
	s.router.GET("/v1/event/GetRecentEvents", c.GetRecentEvents)
}

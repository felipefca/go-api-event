package services

import (
	"context"
	"go-api-event/internal/db/redisdb"
	"go-api-event/internal/models"
	rabbitmq "go-api-event/internal/rabbitMQ"
)

type Service interface {
	PublishEvent(ctx context.Context, message models.Message) error
	GetRecentEvents(ctx context.Context) ([]models.Event, error)
}

type service struct {
	rabbitService rabbitmq.RabbitMQService
	eventRedisDB  redisdb.EventRedisDB
}

func NewService(rabbitService rabbitmq.RabbitMQService, eventRedisDB redisdb.EventRedisDB) Service {
	return &service{
		rabbitService: rabbitService,
		eventRedisDB:  eventRedisDB,
	}
}

func (s service) PublishEvent(ctx context.Context, msg models.Message) error {
	messageBytes := []byte(msg.Message)
	err := s.rabbitService.SendMessage(ctx, messageBytes)
	if err != nil {
		return err
	}

	return nil
}

func (s service) GetRecentEvents(ctx context.Context) ([]models.Event, error) {
	events, err := s.eventRedisDB.GetRecentEvents(ctx)
	if err != nil {
		return nil, err
	}

	return events, nil
}

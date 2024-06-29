package main

import (
	"context"
	"fmt"
	"go-api-event/configs"
	"go-api-event/internal/appctx"
	"go-api-event/internal/server"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	logConfig := zap.NewProductionConfig()

	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	ctx = appctx.WithLogger(ctx, logger)

	amqpConn, err := connectRabbitMQ(ctx)
	if err != nil {
		panic(err)
	}

	defer amqpConn.Close()

	logger.Info("RabbitMQ Connected!")

	s := server.NewServer(server.ServerOptions{
		Logger:   logger,
		Context:  ctx,
		AmqpConn: amqpConn,
	})
	s.Start()
}

func connectRabbitMQ(ctx context.Context) (*amqp.Connection, error) {
	logger := appctx.FromContext(ctx)
	countRetry := 0

	for {
		cfg := configs.GetConfig().RabbitMQ
		amqpUri := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", cfg.UserName, cfg.Password, cfg.HostName, cfg.Port, cfg.VirtualHost)

		conn, err := amqp.Dial(amqpUri)
		if err != nil {
			countRetry++
		} else {
			return conn, nil
		}

		if countRetry <= int(cfg.MaxRetry) {
			logger.Error(fmt.Sprintf("fail to connect RabbitMQ. Retry %d%d...", countRetry, cfg.MaxRetry))
			continue
		} else {
			return nil, fmt.Errorf("error to connect RabbitMQ: %w", err)
		}
	}
}

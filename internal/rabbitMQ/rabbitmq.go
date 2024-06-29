package rabbitmq

import (
	"context"
	"go-api-event/configs"
	"go-api-event/internal/appctx"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQService interface {
	//Send event message to queue
	SendMessage(ctx context.Context, message []byte) error
}

type rabbitMQService struct {
	channel *amqp.Channel
	cfg     configs.RabbitMQ
}

func NewRabbitMQService(conn *amqp.Connection, cfg configs.RabbitMQ) (*rabbitMQService, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	// defer channel.Close()
	// defer conn.Close()

	messeger := rabbitMQService{
		channel: channel,
		cfg:     cfg,
	}

	err = messeger.setup()
	if err != nil {
		return nil, err
	}

	return &messeger, nil
}

func (r *rabbitMQService) SendMessage(ctx context.Context, message []byte) error {
	logger := appctx.FromContext(ctx)

	err := r.channel.PublishWithContext(
		ctx,
		r.cfg.ExchangeName,
		r.cfg.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         message,
			DeliveryMode: amqp.Persistent,
			Headers: amqp.Table{
				"x-correlation-id": appctx.GetCorrelationId(ctx),
			},
		},
	)
	if err != nil {
		logger.Error("error sending message to queue!")
		return err
	}

	return nil
}

func (r *rabbitMQService) setup() error {
	if err := r.buildDeadLetterQueue(); err != nil {
		return err
	}

	if err := r.buildQueue(); err != nil {
		return err
	}

	if err := r.buildErrorQueue(); err != nil {
		return err
	}

	return nil
}

func (r *rabbitMQService) buildDeadLetterQueue() error {
	//Dead Letter Exchange
	err := r.channel.ExchangeDeclare(
		r.cfg.DeadLetterExchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	//Dead Letter Queue
	_, err = r.channel.QueueDeclare(
		r.cfg.DeadLetterQueueName,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    r.cfg.ExchangeName,
			"x-dead-letter-routing-key": r.cfg.RoutingKey,
			"x-message-ttl":             r.cfg.DeadLetterTTL,
		},
	)
	if err != nil {
		return err
	}

	//Bind Dead Letter Queue to Dead Letter Exchange
	r.channel.QueueBind(r.cfg.DeadLetterQueueName, "", r.cfg.DeadLetterExchangeName, false, nil)
	return nil
}

func (r *rabbitMQService) buildQueue() error {
	//Exchange
	err := r.channel.ExchangeDeclare(
		r.cfg.ExchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	//Queue
	_, err = r.channel.QueueDeclare(
		r.cfg.QueueName,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange": r.cfg.DeadLetterExchangeName,
		},
	)
	if err != nil {
		return err
	}

	//Bind Queue to Exchange
	r.channel.QueueBind(r.cfg.QueueName, r.cfg.RoutingKey, r.cfg.ExchangeName, false, nil)
	return nil
}

func (r *rabbitMQService) buildErrorQueue() error {
	//Error Queue
	_, err := r.channel.QueueDeclare(
		r.cfg.ErrorQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

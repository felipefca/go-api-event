package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	Server   Server
	RabbitMQ RabbitMQ
}

type Server struct {
	Port string
}

type RabbitMQ struct {
	HostName               string
	VirtualHost            string
	Port                   int32
	UserName               string
	Password               string
	DeadLetterExchangeName string
	DeadLetterQueueName    string
	DeadLetterTTL          int32
	ExchangeName           string
	QueueName              string
	ErrorQueueName         string
	RoutingKey             string
	MaxRetry               int32
	PrefetchCount          int32
}

func GetConfig() config {
	return *cfg
}

func init() {
	viper.SetDefault("PORT", "8080")

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	viper.ReadInConfig()

	cfg = &config{
		Server: Server{
			Port: viper.GetString("PORT"),
		},
		RabbitMQ: RabbitMQ{
			HostName:               viper.GetString("RABBITMQ_HOST"),
			VirtualHost:            viper.GetString("RABBITMQ_VHOST"),
			Port:                   viper.GetInt32("RABBITMQ_PORT"),
			UserName:               viper.GetString("RABBITMQ_USER"),
			Password:               viper.GetString("RABBITMQ_PASS"),
			DeadLetterExchangeName: viper.GetString("RABBITMQ_DLQ_EXCHANGE"),
			DeadLetterQueueName:    viper.GetString("RABBITMQ_DLQ_QUEUE"),
			DeadLetterTTL:          viper.GetInt32("RABBITMQ_TTL"),
			ExchangeName:           viper.GetString("RABBITMQ_EXCHANGE"),
			QueueName:              viper.GetString("RABBITMQ_QUEUE"),
			ErrorQueueName:         viper.GetString("RABBITMQ_ERROR_QUEUE"),
			RoutingKey:             viper.GetString("RABBITMQ_ROUTING_KEY"),
			MaxRetry:               viper.GetInt32("RABBITMQ_MAX_RETRY"),
			PrefetchCount:          viper.GetInt32("RABBITMQ_PREFETCH_COUNT"),
		},
	}
}

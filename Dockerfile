# Imagem base
FROM golang:1.22.4

# Diretório de trabalho
WORKDIR /app

# Copiar arquivos go.mod e go.sum primeiro, depois instalar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante dos arquivos da aplicação
COPY . .

# Compilar o projeto
RUN go build -o main ./cmd/main.go

# Expõe a porta do servidor
EXPOSE 8080

# Comando de inicialização do container
CMD ["./main"]

# Configurações de ambiente
ENV API_PORT 8080

# Configurações para Redis
ENV REDIS_PORT 6379
ENV REDIS_CONNECTRETRY 3
ENV REDIS_HOST redis
ENV REDIS_TIMEOUT 5000
ENV REDIS_PASSWORD teste

# Configurações para RabbitMQ
ENV RABBITMQ_USER guest
ENV RABBITMQ_PASS guest
ENV RABBITMQ_HOST localhost
ENV RABBITMQ_PORT 5672
ENV RABBITMQ_VHOST teste
ENV RABBITMQ_QUEUE event-queue
ENV RABBITMQ_EXCHANGE event-exchange
ENV RABBITMQ_ROUTING_KEY event-routing
ENV RABBITMQ_DLQ_EXCHANGE event-exchange-dlq
ENV RABBITMQ_DLQ_QUEUE event-queue-dlq
ENV RABBITMQ_ERROR_QUEUE event-error-queue
ENV RABBITMQ_TTL 5000
ENV RABBITMQ_MAX_RETRY 3
ENV RABBITMQ_PREFETCH_COUNT 10
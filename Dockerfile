# Imagem base
FROM golang:1.22.4

# Diretório de trabalho
WORKDIR /app

# Copiar arquivos necessários para o container
COPY . .

# Compilar o projeto
RUN go build -o main .

# Expõe a porta do servidor
EXPOSE 8080

# Comando de inicialização do container
CMD [".cmd/main"]


ENV API_PORT 8080

# Configurações para MongoDB
# ENV CONNECTION_STRING mongodb://teste123:e296cd9f@localhost:27017/admin?authSource=admin
# ENV STOCK_COLLECTION STOCKS
# ENV FIAT_COLLECTION FIATS
# ENV DATABASE CUSTOM_COMMAND

# Configurações para RabbitMQ
ENV RABBITMQ_USER guest
ENV RABBITMQ_PASS guest
# ENV RABBITMQ_HOST http://localhost:15672/
ENV RABBITMQ_HOST http://localhost
ENV RABBITMQ_PORT=15672
ENV RABBITMQ_VHOST=teste
ENV RABBITMQ_QUEUE=event-queue
ENV RABBITMQ_EXCHANGE=event-exchange
ENV RABBITMQ_ROUTING_KEY event-routing
ENV RABBITMQ_DLQ_EXCHANGE event-exchange-dlq
ENV RABBITMQ_DLQ_QUEUE event-queue-dlq
ENV RABBITMQ_ERROR_QUEUE=event-error-queue
ENV RABBITMQ_TTL 5000
ENV RABBITMQ_MAX_RETRY=3
ENV RABBITMQ_PREFETCH_COUNT=10
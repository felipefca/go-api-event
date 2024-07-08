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
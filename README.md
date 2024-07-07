# go-api-event
API in Go with GIN that publish events in RabbitMQ and query in Redis/MongoDB.

Visão Geral
Esta aplicação é uma API em Go que utiliza gin-gonic. Publica eventos no RabbitMQ e faz query no Redis

Pré-requisitos
Certifique-se de ter os seguintes itens instalados antes de prosseguir:

Docker
Docker Compose
Instalação e Execução
Para rodar a aplicação localmente, siga estes passos:

Clone este repositório:

bash
Copiar código
git clone https://github.com/felipefca/go-api-event.git
cd seu-repositorio
Inicie o ambiente usando Docker Compose:

bash
Copiar código
docker-compose up
Isso iniciará a aplicação junto com os serviços necessários.

Para parar a aplicação e remover os containers, use:

bash
Copiar código
docker-compose down
Estrutura do Projeto
Explicação sucinta da estrutura de diretórios da sua aplicação.

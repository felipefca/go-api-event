# go-api-event
This application is a Go API that uses gin-gonic. It publishes events to RabbitMQ and queries Redis

[![LinkedIn][linkedin-shield]][linkedin-url]


<!-- SOBRE O PROJETO -->
## Sobre o Projeto

This API has two endpoints: one for publishing events to an exchange in RabbitMQ and another for querying the latest events in Redi

![img](https://user-images.githubusercontent.com/21323326/233877399-487d793c-76b4-445b-88fd-111c94145c26.png)

### Utilizando

* [![Go][Go-badge]][Go-url]
* [![Redis](https://img.shields.io/badge/Redis-v6.0-red.svg)](https://redis.io/)
* [![RabbitMQ][RabbitMQ-badge]][RabbitMQ-url]
* [![Docker][Docker-badge]][Docker-url]

<!-- GETTING STARTED -->
## Getting Started

Instructions for running the application

### Prerequisites

Run the command to initialize MongoDB, RabbitMQ, and the application on the selected port
* docker
  ```sh
  docker-compose up -d
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/felipefca/go-api-event.git
   ```

2. Exec
   ```sh
   go mod tidy
   ```
   
   ```sh
   cp .env.example .env
   ```
      
   ```sh
   go run ./cmd/main.go
   ```


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/felipe-fernandes-fca/
[Go-url]: https://golang.org/
[Go-badge]: https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white
[RabbitMQ-badge]: https://img.shields.io/badge/rabbitmq-%23ff6600.svg?style=flat&logo=rabbitmq&logoColor=white
[RabbitMQ-url]: https://www.rabbitmq.com/
[Docker-badge]: https://img.shields.io/badge/docker-%230db7ed.svg?style=flat&logo=docker&logoColor=white
[Docker-url]: https://www.docker.com/

package server

import (
	"context"
	"fmt"
	"go-api-event/configs"
	"go-api-event/internal/middlewares"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	amqp "github.com/rabbitmq/amqp091-go"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Start()
}

type ServerOptions struct {
	Logger    *zap.Logger
	Context   context.Context
	AmqpConn  *amqp.Connection
	RedisConn *redis.Client
}

type server struct {
	router *gin.Engine
	ServerOptions
}

func NewServer(opt ServerOptions) Server {
	return server{
		router:        gin.New(),
		ServerOptions: opt,
	}
}

func (s server) Start() {
	s.setupSwagger()
	s.setupServer()
	s.setupMiddlewares()
	s.registerRoutes()

	s.start()
}

func (s server) setupServer() {
	gin.DebugPrintRouteFunc = func(httpmethod, absolutePath, _ string, _ int) {
		s.Logger.Info(fmt.Sprintf("Mapped [%v %v] route", httpmethod, absolutePath))
	}

	s.router.SetTrustedProxies(nil)
}

func (s server) setupMiddlewares() {
	s.router.Use(middlewares.CorrelationIdMiddleware())
	s.router.Use(gin.Recovery())
}

func (s server) setupSwagger() {
	s.router.StaticFile("/swagger.json", "./api/swagger.json")
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s server) start() {
	port := fmt.Sprintf(":%s", configs.GetConfig().Server.Port)

	srv := &http.Server{
		Addr:    port,
		Handler: s.router,
	}

	go func() {
		s.Logger.Info(fmt.Sprintf("Starting server on port %s...", port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Error(fmt.Sprintf("start server error: ", err))
			os.Exit(1)
		}
	}()

	// Aguardar sinal de término
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.Logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		s.Logger.Error(fmt.Sprintf("server shutdown error: ", err))
	}
	s.Logger.Info("Server stopped")
}

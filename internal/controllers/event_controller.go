package controllers

import (
	"go-api-event/internal/appctx"
	"go-api-event/internal/constants"
	"go-api-event/internal/models"
	rabbitmq "go-api-event/internal/rabbitMQ"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EventController struct {
	logger          *zap.Logger
	rabbitMQService rabbitmq.RabbitMQService
}

func NewEventController(logger *zap.Logger, rabbitMQService rabbitmq.RabbitMQService) *EventController {
	return &EventController{
		logger:          logger,
		rabbitMQService: rabbitMQService,
	}
}

// @BasePath /v1

// @Summary Publica um evento de mensagem no RabbitMQ
// @Description Publica uma nova mensagem no RabbitMQ
// @Tags Event
// @Accept json
// @Produce json
// @Param x-correlation-id header string false "Correlation Id"
// @Param message body models.Message true "Mensagem a ser publicada"
// @Success 200 {object} map[string]interface{} "Resposta de sucesso"
// @Router /event/publish [post]
func (e *EventController) Publish(c *gin.Context) {
	ctx := appctx.WithLogger(c, e.logger)
	logger := appctx.FromContext(ctx)

	correlationId := c.Request.Header.Get(constants.CorrelationIdHeader)
	ctx = appctx.SetCorrelationId(ctx, correlationId)

	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageBytes := []byte(msg.Message)
	err := e.rabbitMQService.SendMessage(ctx, messageBytes)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Mensagem publicada com sucesso"})
}

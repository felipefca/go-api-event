package controllers

import (
	"go-api-event/internal/appctx"
	"go-api-event/internal/constants"
	"go-api-event/internal/models"
	"go-api-event/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EventController struct {
	logger  *zap.Logger
	service services.Service
}

func NewEventController(logger *zap.Logger, service services.Service) *EventController {
	return &EventController{
		logger:  logger,
		service: service,
	}
}

// @BasePath /

// @Summary Publica um evento de mensagem no RabbitMQ
// @Description Publica uma nova mensagem no RabbitMQ
// @Tags Event
// @Accept json
// @Produce json
// @Param x-correlation-id header string false "Correlation Id"
// @Param message body models.Message true "Mensagem a ser publicada"
// @Success 200 {object} map[string]interface{} "Resposta de sucesso"
// @Router /v1//event/publish [post]
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

	err := e.service.PublishEvent(ctx, msg)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Mensagem publicada com sucesso"})
}

// @BasePath /

// @Summary Retorna todos os eventos
// @Description Retorna todos os eventos dos últimos 5 minutos
// @Tags Event
// @Accept json
// @Produce json
// @Success 200 {object} []models.Event
// @Router /v1/event/GetRecentEvents [get]
func (e *EventController) GetRecentEvents(c *gin.Context) {
	ctx := appctx.WithLogger(c, e.logger)
	logger := appctx.FromContext(ctx)

	correlationId := c.Request.Header.Get(constants.CorrelationIdHeader)
	ctx = appctx.SetCorrelationId(ctx, correlationId)

	resp, err := e.service.GetRecentEvents(ctx)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

package middlewares

import (
	"go-api-event/internal/constants"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CorrelationIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationId := c.Request.Header.Get(constants.CorrelationIdHeader)
		if strings.TrimSpace(correlationId) == "" {
			correlationId = uuid.NewString()
			c.Request.Header.Add(constants.CorrelationIdHeader, correlationId)
		}
	}
}

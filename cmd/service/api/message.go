package api

import (
	"twitch_chat_analysis/cmd/service/api/schema"
	"twitch_chat_analysis/domain/services"
	"twitch_chat_analysis/extensions/telemetrygt"

	"github.com/gin-gonic/gin"
)

type PushMessageHandler struct {
	PushMessageService *services.PushMessageService
}

func (p PushMessageHandler) pushMessage(c *gin.Context) {
	var message schema.Message

	if err := c.BindJSON(&message); err != nil {
		telemetrygt.Info("Failed to bind JSON")

		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	telemetrygt.Info("Received message", message)

	// Push message to RabbitMQ
	err := p.PushMessageService.PushMessage(message)
	if err != nil {
		telemetrygt.Error("Failed to push message to RabbitMQ", err)

		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.Status(200)
}

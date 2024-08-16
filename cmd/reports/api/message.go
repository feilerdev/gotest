package api

import (
	"twitch_chat_analysis/domain/services"

	"github.com/gin-gonic/gin"
)

type ListMessageHandler struct {
	ListMessageService *services.ListMessageService
}

func (p ListMessageHandler) listMessage(c *gin.Context) {
	// TODO get params
	listInput := services.ListMessageInput{}
	messages := p.ListMessageService.ListMessages(listInput)

	c.IndentedJSON(200, messages)
}

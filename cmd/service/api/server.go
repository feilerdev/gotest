package api

import (
	"twitch_chat_analysis/domain/services"

	"github.com/gin-gonic/gin"
)

func NewServer(pushMessageService *services.PushMessageService) *gin.Engine {
	PushMessageHandler := PushMessageHandler{
		PushMessageService: pushMessageService,
	}

	router := gin.Default()
	router.POST("/message", PushMessageHandler.pushMessage)

	return router
}

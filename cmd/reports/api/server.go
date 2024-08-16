package api

import (
	"twitch_chat_analysis/domain/services"

	"github.com/gin-gonic/gin"
)

func NewServer(listMessageService *services.ListMessageService) *gin.Engine {
	PushMessageHandler := ListMessageHandler{
		ListMessageService: listMessageService,
	}

	router := gin.Default()
	router.GET("/message/list", ListMessageHandler.listMessage)

	return router
}

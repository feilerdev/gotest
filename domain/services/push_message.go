package services

import "twitch_chat_analysis/gateways/messaging/message"

type PushMessageService struct {
	PushMessageService *message.PushMessageService
}

type PushMessageInput struct {
	Message string
}

func (p PushMessageService) PushMessage(input PushMessageInput) error {
	return p.PushMessageService.PushMessage(input.Message)
}

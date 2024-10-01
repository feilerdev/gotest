package services

import "twitch_chat_analysis/domain/vos"

type ListMessageService struct {
}

// TODO: test new
type ListMessageInput struct {
	Message string
}

type ListMessageOutput struct {
	Messages []vos.Message
}

func (l ListMessageService) ListMessages(input ListMessageInput) ListMessageOutput {
	return ListMessageOutput{}
}

package message

import "github.com/wagslane/go-rabbitmq"

type PushMessageService struct {
	Publisher *rabbitmq.Publisher
}

func (p PushMessageService) PushMessage(message string) error {
	return p.Publisher.Publish(
		[]byte(message),
		[]string{"my_routing_key"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("events"),
	)
}

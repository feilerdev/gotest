package messaging

import (
	"log"

	"github.com/wagslane/go-rabbitmq"
)

func NewPublisher(conn *rabbitmq.Conn) *rabbitmq.Publisher {
	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("events"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}

	return publisher
}

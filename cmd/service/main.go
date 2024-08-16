package main

import (
	"twitch_chat_analysis/cmd/service/api"
	"twitch_chat_analysis/domain/services"
	"twitch_chat_analysis/extensions/telemetrygt"
	"twitch_chat_analysis/gateways/messaging/message"

	"github.com/wagslane/go-rabbitmq"
)

func main() {
	logger := telemetrygt.NewLogger()
	logger.Println("Starting service")

	conn, err := rabbitmq.NewConn(
		"amqp://rabbit:pass@rabbitmq",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	publisher, err := rabbitmq.NewPublisher(conn)
	if err != nil {
		logger.Fatal(err)
	}

	pushMessageService := &services.PushMessageService{
		PushMessageService: &message.PushMessageService{
			Publisher: publisher,
		},
	}

	server := api.NewServer(pushMessageService)
	server.Run("localhost:8088")
}

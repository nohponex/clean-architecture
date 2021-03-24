package main

import (
	"context"
	"github.com/nohponex/clean-architecture/internal/notification/adapters"
	"github.com/nohponex/clean-architecture/internal/notification/application"
	"github.com/nohponex/clean-architecture/internal/simplebank/infrastructure"
	"github.com/nohponex/clean-architecture/internal/simplebank/infrastructure/rabbitmq"
)

func main() {
	ctx := context.Background()

	config, err := infrastructure.NewConfigFromEnvironmental()
	if err != nil {
		panic(err)
	}

	connection, err := rabbitmq.NewConnectionFromConfig(config)
	if err != nil {
		panic(err)
	}

	if err := adapters.NewRabbitMQConsumer(
		ctx,
		connection,
		[]application.Notification{
			application.NewEmailNotification(),
			application.NewSMSNotification(),
		},
	); err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"fmt"
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

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	queue := "notification"
	if _, err := channel.QueueDeclare(
		queue, // name of the queue
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	); err != nil {
		panic(err)
	}

	//bind

	deliveries, err := channel.Consume(
		queue,         // name
		queue+"-ABCD", // consumerTag,
		false,         // noAck
		false,         // exclusive
		false,         // noLocal
		false,         // noWait
		nil,           // arguments
	)
	if err != nil {
		panic(err)
	}

	if err := channel.QueueBind(
		queue,        // name of the queue
		"#",          // bindingKey
		"simplebank", // sourceExchange
		false,        // noWait
		nil,          // arguments
	); err != nil {
		panic(err)
	}

	emailNotification := application.NewEmailNotification()
	withdrawnAdapter := adapters.NewWithdrawnAdapter(emailNotification)

	for {
		select {
		case d, ok := <-deliveries:
			if !ok {
				break
			}

			err := withdrawnAdapter.Handle(ctx, d)
			if err != nil {
				fmt.Println(err.Error())
				_ = d.Nack(false, true)
				break
			}

			_ = d.Ack(false)
		}
	}
}

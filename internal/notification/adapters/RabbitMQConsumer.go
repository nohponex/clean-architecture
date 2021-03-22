package adapters

import (
	"context"
	"fmt"
	"github.com/nohponex/clean-architecture/internal/notification/application"
	"github.com/streadway/amqp"
	"math/rand"
	"time"
)

func NewRabbitMQConsumer(
	ctx context.Context,
	connection *amqp.Connection,
	useCase application.Notification,
) error {
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	adapter := NewWithdrawnAdapter(useCase)
	handle(ctx, channel, "notification", adapter)

	return nil
}

func handle(ctx context.Context, channel *amqp.Channel, queue string, adapter RabbitMQAdapter) {
	const exchange = "simplebank"

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

	if err := channel.QueueBind(
		queue,    // name of the queue
		"#",      // bindingKey
		exchange, // sourceExchange
		false,    // noWait
		nil,      // arguments
	); err != nil {
		panic(err)
	}

	rand.Seed(time.Now().Unix())

	deliveries, err := channel.Consume(
		queue, // name
		fmt.Sprintf("%s-consumer-%d", queue, rand.Int31()),
		false, // noAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case d, ok := <-deliveries:
			if !ok {
				break
			}

			err := adapter.Handle(ctx, d)
			if err != nil {
				fmt.Println(err.Error())
				_ = d.Nack(false, true)
				break
			}

			_ = d.Ack(false)
		}
	}
}

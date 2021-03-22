package adapters

import (
	"context"
	"fmt"
	"github.com/nohponex/clean-architecture/internal/notification/application"
	"github.com/streadway/amqp"
	"math/rand"
	"reflect"
	"time"
)

func NewRabbitMQConsumer(
	ctx context.Context,
	connection *amqp.Connection,
	notificationUseCases []application.Notification,
) error {
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	for _, useCase := range notificationUseCases {
		useCaseReflection := reflect.TypeOf(useCase)
		adapter := NewWithdrawnAdapter(useCase)
		go handle(ctx, channel, "notification-"+useCaseReflection.String(), adapter)
	}
	rand.Seed(time.Now().Unix())

	<-ctx.Done()
	return nil
}

func handle(ctx context.Context, channel *amqp.Channel, queue string, adapter RabbitMQAdapter) {
	const exchange = "simplebank"

	if _, err := channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		panic(err)
	}

	if err := channel.QueueBind(
		queue,
		"#",
		exchange,
		false,
		nil,
	); err != nil {
		panic(err)
	}

	deliveries, err := channel.Consume(
		queue,
		fmt.Sprintf("%s-consumer-%d", queue, rand.Int31()),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("consuming from " + queue)
	for {
		select {
		case d, ok := <-deliveries:
			if !ok {
				break
			}

			fmt.Println("reading " + string(d.Body))
			err := adapter.Handle(ctx, d)
			if err != nil {
				fmt.Println(err.Error())
				if d.Redelivered {
					_ = d.Nack(false, false)
				} else {
					_ = d.Nack(false, true)
				}

				break
			}

			_ = d.Ack(false)
		}
	}
}

package messaging

import (
	"context"
	"encoding/json"
	"github.com/nohponex/clean-architecture/internal/simplebank/domain/repositories"
	"github.com/nohponex/clean-architecture/internal/simplebank/infrastructure"
	"github.com/nohponex/clean-architecture/internal/simplebank/infrastructure/rabbitmq"
	"github.com/streadway/amqp"
)

type rabbitMQMessaging struct {
	channel *amqp.Channel
}

func NewRabbitMQMessaging(config *infrastructure.Config) (repositories.Messaging, error) {
	connection, err := rabbitmq.NewConnectionFromConfig(config)
	if err != nil {
		return nil, err
	}
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &rabbitMQMessaging{channel: channel}, nil
}

func (a rabbitMQMessaging) Publish(
	ctx context.Context,
	eventType string,
	body map[string]interface{},
) error {
	message, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return a.channel.Publish(
		"simplebank",
		eventType,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         message,
			DeliveryMode: amqp.Persistent,
		},
	)
}

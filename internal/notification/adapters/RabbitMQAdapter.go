package adapters

import (
	"context"
	"github.com/streadway/amqp"
)

type RabbitMQAdapter interface {
	Handle(ctx context.Context, delivery amqp.Delivery) error
}

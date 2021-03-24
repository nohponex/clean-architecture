package rabbitmq

import (
	"fmt"
	"github.com/nohponex/clean-architecture/internal/simplebank/infrastructure"
	"github.com/streadway/amqp"
)

func NewConnectionFromConfig(
	config *infrastructure.Config,
) (
	*amqp.Connection,
	error,
) {
	uri := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		config.AMQ_USERNAME,
		config.AMQ_PASSWORD,
		config.AMQ_HOST,
		config.AMQ_PORT,
		config.AMQ_VHOST,
	)

	return amqp.Dial(uri)
}

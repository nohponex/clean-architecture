package adapters

import (
	"context"
	"encoding/json"
	"github.com/nohponex/clean-architecture/internal/notification/application"
	"github.com/streadway/amqp"
)

type AccountWithdrawn struct {
	AccountID string  `json:"AccountID"`
	PersonID  string  `json:"PersonID"`
	Amount    float32 `json:"Amount"`
	Currency  string  `json:"Currency"`
}

type WithdrawnAdapter struct {
	notification application.Notification
}

func NewWithdrawnAdapter(notification application.Notification) RabbitMQAdapter {
	return &WithdrawnAdapter{notification: notification}
}

func (a *WithdrawnAdapter) Handle(ctx context.Context, delivery amqp.Delivery) error {
	event := AccountWithdrawn{}
	if err := json.Unmarshal(delivery.Body, &event); err != nil {
		return err
	}

	return a.notification.Withdrawn(
		ctx,
		event.AccountID,
		event.Amount,
		event.Currency,
	)
}

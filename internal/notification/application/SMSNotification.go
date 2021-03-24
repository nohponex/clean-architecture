package application

import (
	"context"
	"fmt"
)

type smsNotification struct {
}

func NewSMSNotification() Notification {
	return &smsNotification{}
}

func (e smsNotification) Withdrawn(ctx context.Context, accountID string, amount float32, currency string) error {
	fmt.Printf("SMS Notification for %q becase withdrawn %f %s\n", accountID, amount, currency)

	return nil
}

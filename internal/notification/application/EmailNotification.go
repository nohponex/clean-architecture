package application

import (
	"context"
	"fmt"
)

type emailNotification struct {
}

func NewEmailNotification() Notification {
	return &emailNotification{}
}

func (e emailNotification) Withdrawn(ctx context.Context, accountID string, amount float32, currency string) error {
	fmt.Printf("Email Notification for %q becase withdrawn %f %s\n", accountID, amount, currency)

	return nil
}

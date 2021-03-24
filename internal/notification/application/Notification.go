package application

import "context"

type Notification interface {
	Withdrawn(
		ctx context.Context,
		accountID string,
		amount float32,
		currency string,
	) error
}

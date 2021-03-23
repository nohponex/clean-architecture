package repositories

import "context"

type Messaging interface {
	Publish(
		ctx context.Context,
		eventType string,
		body map[string]interface{},
	) error
}

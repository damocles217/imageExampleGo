package pkg

import "context"

type Publisher interface {
	Publish(ctx context.Context, payload interface{}) error
}

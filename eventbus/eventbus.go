package eventbus

import "context"

type SubscribeOption func(any)

type EventBus interface {
	Subscribe(ctx context.Context, eventName string, handler Handler, opts ...SubscribeOption) error
	Publish(ctx context.Context, event Event) error
	Close()
}

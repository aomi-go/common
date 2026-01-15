package adapters

import (
	"context"
	"errors"
	"time"

	"github.com/aomi-go/common/eventbus"
	"github.com/valkey-io/valkey-go"
)

func NewValkeyEventBus(
	parent context.Context,
	client valkey.Client,
) *ValkeyEventBus {
	ctx, cancel := context.WithCancel(parent)

	return &ValkeyEventBus{
		client: client,
		ctx:    ctx,
		cancel: cancel,
	}
}

type ValkeyEventBus struct {
	client valkey.Client
	ctx    context.Context
	cancel context.CancelFunc
}

func (v *ValkeyEventBus) Subscribe(ctx context.Context, eventName string, handler eventbus.Handler, opts ...eventbus.SubscribeOption) error {
	go func() {
		err := v.client.Receive(v.ctx, v.client.B().Subscribe().Channel(eventName).Build(), func(msg valkey.PubSubMessage) {
			// 3. 异步执行 handler 是对的，防止阻塞 valkey 读取流
			go handler(v.ctx, &eventbus.BaseEvent{
				EventName: msg.Channel,
				EventAt:   time.Now(),
				Data:      msg.Message,
			})
		})

		if err != nil && !errors.Is(err, context.Canceled) {
			// 这里应该有日志记录，或者推送到某种错误处理通道
			log.Err(err).Errorf("Valkey subscribe error for %s", eventName)
		}
	}()

	return nil
}

func (v *ValkeyEventBus) Publish(ctx context.Context, event eventbus.Event) error {
	cmd := v.client.B().Publish().Channel(event.GetName()).Message(event.GetData().(string)).Build()
	res := v.client.Do(ctx, cmd)
	return res.Error()
}

func (v *ValkeyEventBus) Close() {
}

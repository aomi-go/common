package valkey

import (
	"context"
	"testing"
	"time"

	"github.com/aomi-go/common/eventbus"
	"github.com/valkey-io/valkey-go"
)

func TestEventBus(t *testing.T) {
	vc, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{"10.43.180.156:6379"}})
	if err != nil {
		t.Fatalf("%s", err)
	}

	eventBus := NewValkeyEventBus(context.TODO(), vc)
	_ = eventBus.Publish(context.TODO(), &eventbus.BaseEvent{
		EventName: "test",
		Data:      "hello world",
	})

	eventBus.Subscribe(context.TODO(), "test", func(ctx context.Context, e eventbus.Event) {
		t.Logf("我是第一个: %s", e.GetData())
	})

	eventBus.Subscribe(context.TODO(), "test", func(ctx context.Context, e eventbus.Event) {
		t.Logf("我是第二个: %s", e.GetData())
	})

	for {
		eventBus.Publish(context.TODO(), &eventbus.BaseEvent{
			EventName: "test",
			Data:      "hello world!",
		})
		time.Sleep(2 * time.Second)
	}
}

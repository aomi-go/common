package memory

import (
	"context"
	"testing"
	"time"

	"github.com/aomi-go/common/eventbus"
)

func TestEventBus(t *testing.T) {

	eventBus := NewMemoryEventBus(context.TODO())
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

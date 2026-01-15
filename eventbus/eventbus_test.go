package eventbus

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestEventBus(t *testing.T) {

	eventBus := NewEventBus(context.TODO())
	eventBus.Publish(testEvent{msg: "hello"})

	eventBus.Subscribe("test", func(ctx context.Context, e Event) {
		te := e.(testEvent)
		t.Logf("我是第一个: %s", te.msg)
	})

	eventBus.Subscribe("test", func(ctx context.Context, e Event) {
		te := e.(testEvent)
		t.Logf("我是第二个: %s", te.msg)
	})

	for {
		eventBus.Publish(testEvent{msg: "hello"})
		time.Sleep(2 * time.Second)
	}
}

func TestEventBus_ConcurrentPublish(t *testing.T) {
	bus := NewEventBus(context.Background())
	defer bus.Close()

	var count int64

	bus.Subscribe(
		"test",
		func(ctx context.Context, e Event) {
			atomic.AddInt64(&count, 1)
		},
		WithWorkers(4),
		WithBuffer(1000),
	)

	const total = 10000

	var wg sync.WaitGroup
	wg.Add(total)

	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			bus.Publish(testEvent{})
		}()
	}

	wg.Wait()
	time.Sleep(200 * time.Millisecond)

	if atomic.LoadInt64(&count) != total {
		t.Fatalf("expected %d, got %d", total, count)
	}
}

func BenchmarkEventBus_Publish(b *testing.B) {
	bus := NewEventBus(context.Background())
	defer bus.Close()

	bus.Subscribe(
		"test",
		func(ctx context.Context, e Event) {
		},
		WithWorkers(runtime.NumCPU()),
		WithBuffer(4096),
	)

	ev := testEvent{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bus.Publish(ev)
	}
}

type testEvent struct {
	msg string
}

func (testEvent) Name() string {
	return "test"
}

func (testEvent) Time() time.Time {
	return time.Now()
}

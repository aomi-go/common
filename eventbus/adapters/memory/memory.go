package memory

import (
	"context"
	"runtime"
	"sync"

	"github.com/aomi-go/common/eventbus"
	"github.com/aomi-go/common/logger"
)

var log = logger.GetLogger("adapter_memory")

type subscribeOptions struct {
	workers int
	buffer  int
}

func defaultOptions() subscribeOptions {
	return subscribeOptions{
		workers: runtime.NumCPU(),
		buffer:  256,
	}
}

func WithWorkers(n int) eventbus.SubscribeOption {
	return func(o any) {
		if opts, ok := o.(*subscribeOptions); ok {
			opts.workers = n
		}
	}
}

// WithBuffer 设置订阅缓冲区大小
func WithBuffer(n int) eventbus.SubscribeOption {
	return func(o any) {
		if opts, ok := o.(*subscribeOptions); ok {
			opts.buffer = n
		}
	}
}

func NewMemoryEventBus(parent context.Context) *EventBus {
	ctx, cancel := context.WithCancel(parent)
	return &EventBus{
		handlers: make(map[string][]*workerPool),
		ctx:      ctx,
		cancel:   cancel,
	}
}

type EventBus struct {
	mu       sync.RWMutex
	handlers map[string][]*workerPool
	ctx      context.Context
	cancel   context.CancelFunc
	closed   bool
}

func (b *EventBus) Subscribe(
	ctx context.Context,
	eventName string,
	handler eventbus.Handler,
	opts ...eventbus.SubscribeOption,
) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	o := defaultOptions()
	for _, opt := range opts {
		opt(&o)
	}

	wp := newWorkerPool(b.ctx, handler, o.workers, o.buffer)
	b.handlers[eventName] = append(b.handlers[eventName], wp)
	return nil
}

func (b *EventBus) Publish(ctx context.Context, event eventbus.Event) error {
	b.mu.RLock()
	pools := append([]*workerPool(nil), b.handlers[event.GetName()]...)
	b.mu.RUnlock()

	for _, wp := range pools {
		select {
		case wp.ch <- event:
		default:
			// ❗ 丢弃 or 记录指标
		}
	}
	return nil
}

func (b *EventBus) Close() {
	b.mu.Lock()
	if b.closed {
		b.mu.Unlock()
		return
	}
	b.closed = true
	b.cancel()

	for _, pools := range b.handlers {
		for _, wp := range pools {
			close(wp.ch)
		}
	}
	b.mu.Unlock()
}

func newWorkerPool(
	ctx context.Context,
	handler eventbus.Handler,
	workers int,
	buffer int,
) *workerPool {

	wp := &workerPool{
		handler: handler,
		ch:      make(chan eventbus.Event, buffer),
		ctx:     ctx,
	}

	for i := 0; i < workers; i++ {
		go wp.run(i)
	}

	return wp
}

type workerPool struct {
	handler eventbus.Handler
	ch      chan eventbus.Event
	ctx     context.Context
}

func (w *workerPool) run(id int) {
	defer func() {
		if r := recover(); r != nil {
			log.Any("recover", r).Error("recover from panic")
		}
	}()
	for {
		select {
		case <-w.ctx.Done():
			return
		case e, ok := <-w.ch:
			if !ok {
				return
			}
			w.handler(w.ctx, e)
		}
	}
}

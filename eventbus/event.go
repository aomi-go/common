package eventbus

import "time"

type Event interface {
	GetName() string
	GetTime() time.Time
	GetData() any
}

type BaseEvent struct {
	EventName string
	EventAt   time.Time
	Data      any
}

func (b *BaseEvent) GetName() string {
	return b.EventName
}

func (b *BaseEvent) GetTime() time.Time {
	return b.EventAt
}

func (b *BaseEvent) GetData() any {
	return b.Data
}

package eventbus

import "time"

type Event interface {
	Name() string
	Time() time.Time
}

type BaseEvent[T interface{}] struct {
	EventName string
	EventAt   time.Time
	Data      T
}

func (e *BaseEvent[T]) Name() string {
	return e.EventName
}
func (e *BaseEvent[T]) Time() time.Time {
	return e.EventAt
}

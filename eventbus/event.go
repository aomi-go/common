package eventbus

import "time"

type Event interface {
	Name() string
	Time() time.Time
}

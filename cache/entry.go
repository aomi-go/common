package cache

import (
	"fmt"
	"time"
)

func NewEntry(key string, ttl time.Duration) Entry {
	return Entry{Key: key, Ttl: ttl}
}

type Entry struct {
	Key string
	Ttl time.Duration
}

func (e Entry) SetKeySuffix(suffix string) Entry {
	e.Key = fmt.Sprintf("%s:%s", e.Key, suffix)
	return e
}

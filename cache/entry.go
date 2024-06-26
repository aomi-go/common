package cache

import (
	"fmt"
	"time"
)

func NewEntry(key string, ttl time.Duration) Entry {
	return Entry{Key: key, Ttl: ttl}
}

type Entry struct {
	Key string        `json:"key" yaml:"key"`
	Ttl time.Duration `json:"ttl" yaml:"ttl"`
}

func (e Entry) SetKeySuffix(suffix string) Entry {
	e.Key = fmt.Sprintf("%s:%s", e.Key, suffix)
	return e
}

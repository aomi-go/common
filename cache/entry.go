package cache

import "time"

// Entry 用于存储缓存值及其过期时间
type Entry struct {
	Value     interface{}
	ExpiresAt time.Time
}

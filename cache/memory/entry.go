package memory

import "time"

// MemoryEntry 用于存储缓存值及其过期时间
type MemoryEntry struct {
	Value     interface{}
	ExpiresAt time.Time
}

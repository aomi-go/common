package memory

import "time"

// Entry MemoryEntry 用于存储缓存值及其过期时间
type Entry struct {
	Value     []byte
	ExpiresAt time.Time
}

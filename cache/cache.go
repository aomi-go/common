package cache

import (
	"context"
	"errors"
	"time"
)

var (
	ValueNotNilPointer = errors.New("value must be a non-nil pointer")
	ErrEntryNotFound   = errors.New("Entry not found")
)

// Cache 缓存服务接口
type Cache interface {
	// Get 根据key获取缓存值
	// @param key string 缓存key
	// @param value interface{} 获取到的缓存值
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Increment(ctx context.Context, key string, value int64, ttl time.Duration) (int64, error)
	IncrementByFloat(ctx context.Context, key string, value float64, ttl time.Duration) (float64, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) bool
}

type MemoryCache Cache
type PersistCache Cache

type RedisCache Cache
type ValKeyCache Cache

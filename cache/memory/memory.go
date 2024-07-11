package memory

import (
	"context"
	"errors"
	"github.com/allegro/bigcache/v3"
	"github.com/aomi-go/common/cache"
	"sync"
	"time"
)

type CacheConfig bigcache.Config

func DefaultConfig(eviction time.Duration) CacheConfig {
	return CacheConfig(bigcache.DefaultConfig(eviction))
}

func NewMemoryCache(config CacheConfig) *Cache {
	provider, initErr := bigcache.New(context.Background(), bigcache.Config(config))
	if initErr != nil {
		panic(initErr)
	}
	return &Cache{
		provider: provider,
		mu:       &sync.Mutex{},
	}
}

/*
Cache 内存缓存实现
*/
type Cache struct {
	provider *bigcache.BigCache

	mu *sync.Mutex
}

func (m *Cache) Get(ctx context.Context, key string, value interface{}) error {
	data, err := m.provider.Get(key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return cache.ErrEntryNotFound
		}
		return err
	}

	var entry Entry
	err = cache.Deserialize(data, &entry)
	if err != nil {
		return err
	}

	if time.Now().After(entry.ExpiresAt) {
		_ = m.provider.Delete(key)
		return cache.ErrEntryNotFound
	}

	err = cache.Deserialize(entry.Value, value)
	if err != nil {
		return err
	}

	return nil
}

func (m *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	valueBuf, err := cache.Serialize(value)
	if err != nil {
		return err
	}
	entry := Entry{
		Value:     valueBuf.Bytes(),
		ExpiresAt: time.Now().Add(ttl),
	}

	data, err := cache.Serialize(entry)
	if err != nil {
		return err
	}

	return m.provider.Set(key, data.Bytes())
}

func (m *Cache) Delete(ctx context.Context, key string) error {
	return m.provider.Delete(key)
}

func (m *Cache) Exists(ctx context.Context, key string) bool {
	_, err := m.provider.Get(key)
	return err == nil
}

func (m *Cache) Increment(ctx context.Context, key string, value int64, ttl time.Duration) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var currentValue int64
	err := m.Get(ctx, key, &currentValue)
	if err != nil {
		if errors.Is(err, cache.ErrEntryNotFound) {
			currentValue = int64(0)
		} else {
			return 0, err
		}
	}

	var newValue = currentValue + value

	err = m.Set(ctx, key, newValue, ttl)
	if err != nil {
		return 0, err
	}

	return newValue, nil
}

func (m *Cache) IncrementByFloat(ctx context.Context, key string, value float64, ttl time.Duration) (float64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var currentValue float64
	err := m.Get(ctx, key, &currentValue)
	if err != nil {
		if errors.Is(err, cache.ErrEntryNotFound) {
			currentValue = float64(0)
		} else {
			return 0, err
		}
	}

	var newValue = currentValue + value

	err = m.Set(ctx, key, newValue, ttl)
	if err != nil {
		return 0, err
	}

	return newValue, nil
}

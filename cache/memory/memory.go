package memory

import (
	"cache"
	"context"
	"errors"
	"github.com/allegro/bigcache/v3"
	"reflect"
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
			return nil
		}
		return err
	}

	var entry cache.Entry
	err = cache.Deserialize(data, &entry)
	if err != nil {
		return err
	}

	if time.Now().After(entry.ExpiresAt) {
		_ = m.provider.Delete(key)
		return bigcache.ErrEntryNotFound
	}

	// 使用反射将 entry.Value 赋值给 value
	val := reflect.ValueOf(value)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return cache.ValueNotNilPointer
	}

	val.Elem().Set(reflect.ValueOf(entry.Value))

	return nil
}

func (m *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	entry := cache.Entry{
		Value:     value,
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
		return 0, err
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
		if errors.Is(err, bigcache.ErrEntryNotFound) {
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

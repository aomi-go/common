package valkey

import (
	"context"

	"github.com/aomi-go/common/cache"
	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/options"

	"time"
)

func NewCache(client *glide.Client) *Cache {
	return &Cache{
		client: client,
	}
}

type Cache struct {
	client *glide.Client
}

func (c *Cache) Get(ctx context.Context, key string, value interface{}) error {
	result, err := c.client.Get(ctx, key)
	if nil != err {
		return err
	}
	if !result.IsNil() {
		return cache.DeserializeWithStr(result.Value(), value)
	}
	return err
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	buf, err := cache.Serialize(value)
	if nil != err {
		return err
	}
	_, err = c.client.SetWithOptions(ctx, key, buf.String(), options.SetOptions{
		Expiry: options.NewExpiryIn(ttl),
	})
	return err
}

func (c *Cache) Increment(ctx context.Context, key string, value int64, ttl time.Duration) (int64, error) {
	v, err := c.client.IncrBy(ctx, key, value)
	if nil != err {
		return 0, err
	}
	if ttl > 0 {
		_, _ = c.client.Expire(ctx, key, ttl)
	}
	return v, err
}

func (c *Cache) IncrementByFloat(ctx context.Context, key string, value float64, ttl time.Duration) (float64, error) {
	v, err := c.client.IncrByFloat(ctx, key, value)
	if nil != err {
		return 0, err
	}
	if ttl > 0 {
		_, _ = c.client.Expire(ctx, key, ttl)
	}
	return v, err
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	_, err := c.client.Del(ctx, []string{key})
	return err
}

func (c *Cache) Exists(ctx context.Context, key string) bool {
	result, err := c.client.Exists(ctx, []string{key})
	return nil == err && result == 1
}

package valkey

import (
	"context"
	"time"

	"github.com/aomi-go/common/cache"
	"github.com/valkey-io/valkey-go"
)

func NewCache(client valkey.Client) *Cache {
	return &Cache{
		client: client,
	}
}

type Cache struct {
	client valkey.Client
}

func (c *Cache) Get(ctx context.Context, key string, value interface{}) error {
	resp := c.client.Do(ctx, c.client.B().Get().Key(key).Build())
	if resp.Error() != nil {
		if valkey.IsValkeyNil(resp.Error()) {
			return cache.ErrEntryNotFound
		}
		return resp.Error()
	}

	return cache.DeserializeWithStr(resp.String(), value)
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	v, err := cache.Serialize(value)
	if nil != err {
		return err
	}

	cmd := c.client.B().
		Set().
		Key(key).
		Value(v.String()).
		Ex(ttl).
		Build()
	_, err = c.client.Do(ctx, cmd).ToAny()
	return err
}

func (c *Cache) Increment(ctx context.Context, key string, value int64, ttl time.Duration) (int64, error) {
	res, err := c.client.Do(ctx, c.client.B().Incrby().Key(key).Increment(value).Build()).AsInt64()
	if err != nil {
		return 0, err
	}

	if ttl > 0 {
		_ = c.client.Do(ctx, c.client.B().Expire().Key(key).Seconds(int64(ttl.Seconds())).Build())
	}
	return res, nil
}

func (c *Cache) IncrementByFloat(ctx context.Context, key string, value float64, ttl time.Duration) (float64, error) {
	res, err := c.client.Do(ctx, c.client.B().Incrbyfloat().Key(key).Increment(value).Build()).AsFloat64()
	if err != nil {
		return 0, err
	}

	if ttl > 0 {
		_ = c.client.Do(ctx, c.client.B().Expire().Key(key).Seconds(int64(ttl.Seconds())).Build())
	}
	return res, nil
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Do(ctx, c.client.B().Del().Key(key).Build()).Error()
}

func (c *Cache) Exists(ctx context.Context, key string) bool {
	res, err := c.client.Do(ctx, c.client.B().Exists().Key(key).Build()).AsInt64()
	if err != nil {
		return false
	}
	return res > 0
}

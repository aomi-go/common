package redis

import (
	"context"
	"github.com/aomi-go/common/cache"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisCache(client *redis.Client) *Cache {
	return &Cache{
		client: client,
	}
}

var incrByScript = redis.NewScript(`
    local newVal = redis.call('INCRBY', KEYS[1], ARGV[1])
    redis.call('EXPIRE', KEYS[1], ARGV[2])
    return newVal
`)

var incrByFloatScript = redis.NewScript(`
    local newVal = redis.call('INCRBYFLOAT', KEYS[1], ARGV[1])
    redis.call('EXPIRE', KEYS[1], ARGV[2])
    return newVal
`)

type Cache struct {
	client *redis.Client
}

func (c *Cache) Get(ctx context.Context, key string, value interface{}) error {
	str, err := c.client.Get(ctx, key).Result()
	if nil != err {
		return err
	}
	err = cache.DeserializeWithStr(str, value)
	return err
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	buf, err := cache.Serialize(value)
	if nil != err {
		return err
	}

	return c.client.Set(ctx, key, buf.String(), ttl).Err()
}

func (c *Cache) Increment(ctx context.Context, key string, value int64, ttl time.Duration) (int64, error) {
	return incrByScript.Run(ctx, c.client, []string{key}, value, int64(ttl.Seconds())).Int64()
}
func (c *Cache) IncrementByFloat(ctx context.Context, key string, value float64, ttl time.Duration) (float64, error) {
	return incrByFloatScript.Run(ctx, c.client, []string{key}, value, int64(ttl.Seconds())).Float64()
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *Cache) Exists(ctx context.Context, key string) bool {
	result, _ := c.client.Exists(ctx, key).Result()
	return result == 1
}

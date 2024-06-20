package redis

import (
	"cache"
	"context"
	"errors"
	"fmt"
	"github.com/redis/rueidis"
	"strconv"
	"time"
)

func NewRueidisCache(client rueidis.Client) *RueidisCache {
	return &RueidisCache{client: client}
}

type RueidisCache struct {
	client rueidis.Client
}

func (r *RueidisCache) Get(ctx context.Context, key string, value interface{}) error {
	cmd := r.client.B().Get().Key(key).Build()
	resp, err := r.client.Do(ctx, cmd).ToString()
	if err != nil {
		if rueidis.IsRedisNil(err) {
			return nil
		}
		return err
	}
	err = cache.DeserializeWithStr(resp, value)
	return err
}

func (r *RueidisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	v, err := cache.Serialize(value)
	if nil != err {
		return err
	}

	cmd := r.client.B().
		Set().
		Key(key).
		Value(v.String()).
		Ex(ttl).
		Build()
	_, err = r.client.Do(ctx, cmd).ToAny()
	return err
}

func (r *RueidisCache) Increment(ctx context.Context, key string, value int64, ttl time.Duration) (int64, error) {
	cmd := r.client.B().Incrby().Key(key).Increment(value).Build()
	resp, err := r.client.Do(ctx, cmd).ToInt64()
	if err != nil {
		return 0, err
	}
	if ttl > 0 {
		ttlCmd := r.client.B().Expire().Key(key).Seconds(int64(ttl.Seconds())).Build()
		return r.client.Do(ctx, ttlCmd).ToInt64()
	}
	return resp, nil
}

func (r *RueidisCache) IncrementByFloat(ctx context.Context, key string, value float64, ttl time.Duration) (float64, error) {
	cmd := r.client.B().Incrbyfloat().Key(key).Increment(value).Build()
	resp, err := r.client.Do(ctx, cmd).ToAny()
	fmt.Println(resp)
	if err != nil {
		return 0, err
	}
	if ttl > 0 {
		ttlCmd := r.client.B().Expire().Key(key).Seconds(int64(ttl.Seconds())).Build()
		_, err := r.client.Do(ctx, ttlCmd).ToAny()
		if err != nil {
			return 0, err
		}
	}
	str, ok := resp.(string)
	if ok {
		return strconv.ParseFloat(str, 64)
	}
	return 0, errors.New("increment by float failed")
}

func (r *RueidisCache) Delete(ctx context.Context, key string) error {
	cmd := r.client.B().Del().Key(key).Build()
	_, err := r.client.Do(ctx, cmd).ToInt64()
	return err
}

func (r *RueidisCache) Exists(ctx context.Context, key string) bool {
	cmd := r.client.B().Exists().Key(key).Build()
	resp, err := r.client.Do(ctx, cmd).ToInt64()
	if err != nil {
		return false
	}
	return resp > 0
}

package test

import (
	"cache"
	r "cache/redis"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/redis/rueidis"
	"testing"
	"time"
)

type User struct {
	Name     string
	Age      int
	CreateAt time.Time
}

func testCache(c cache.Cache) {
	var v User
	err := c.Set(context.Background(), "struct", User{
		Name:     "hello",
		Age:      18,
		CreateAt: time.Time{},
	}, 1*time.Minute)
	if nil != err {
		fmt.Println(v)
	}
	_ = c.Get(context.Background(), "struct", &v)
	fmt.Println(v)

	var str string
	err = c.Set(context.Background(), "str", "hello world", 1*time.Minute)
	fmt.Println(str)

	var number int64
	err = c.Set(context.Background(), "number", int64(100), 1*time.Minute)
	_ = c.Get(context.Background(), "number", &number)
	fmt.Println(number)

	_, _ = c.Increment(context.Background(), "count", 1, 1*time.Minute)
	_, _ = c.IncrementByFloat(context.Background(), "countByFloat", 1.5, 1*time.Minute)

}

func TestRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	c := r.NewRedisCache(client)
	testCache(c)
}

func TestRueidis(t *testing.T) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	c := r.NewRueidisCache(client)
	testCache(c)
}

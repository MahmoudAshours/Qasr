package redisCache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(addr string) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr, // "localhost:6379"
	})
	return &RedisCache{Client: rdb}
}

func (r *RedisCache) Get(slug string) (string, error) {
	val, err := r.Client.Get(context.Background(), slug).Result()
	if err == redis.Nil {
		return "", nil // cache miss
	}
	return val, err
}

func (r *RedisCache) Set(slug, url string, ttl time.Duration) error {
	return r.Client.Set(context.Background(), slug, url, ttl).Err()
}

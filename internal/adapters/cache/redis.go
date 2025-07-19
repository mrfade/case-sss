package cache

import (
	"context"
	"time"

	"github.com/mrfade/case-sss/internal/adapters/storage/redis"
)

type RedisCache struct {
	client *redis.Storage
}

func NewRedisCache(client *redis.Storage) *RedisCache {
	return &RedisCache{
		client,
	}
}

func (c *RedisCache) Set(key string, value any, expiration time.Duration) error {
	return c.client.Client.Set(context.Background(), key, value, expiration).Err()
}

func (c *RedisCache) Get(key string) (any, error) {
	val, err := c.client.Client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (c *RedisCache) Del(key string) error {
	return c.client.Client.Del(context.Background(), key).Err()
}

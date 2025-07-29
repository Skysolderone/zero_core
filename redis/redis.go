package redis

import (
	"github.com/Skysolderone/zero_core/config"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	redis.Client
}

func NewRedisClient(c *config.Config) *RedisClient {
	return &RedisClient{
		Client: *redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
			DB:       0,
		}),
	}
}

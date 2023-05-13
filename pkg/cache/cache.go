package cache

import (
	"time"

	"context"

	"github.com/je-martinez/2023-go-rest-api/config"

	"github.com/redis/go-redis/v9"
)

func New(cfg *config.RedisConfig) *RedisApiInstance {
	return &RedisApiInstance{
		client: redis.NewClient(&redis.Options{
			Addr:         cfg.RedisAddr,
			MinIdleConns: cfg.MinIdleConns,
			PoolSize:     cfg.PoolSize,
			PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
			Password:     cfg.Password, // no password set
			DB:           cfg.DB,       // use default DB
		}),
	}
}

type RedisApiInstance struct {
	client *redis.Client
}

func (r *RedisApiInstance) Add(ctx context.Context, key string, value interface{}) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisApiInstance) Remove(ctx context.Context, key ...string) (int64, error) {
	return r.client.Del(ctx, key...).Result()
}

func (r *RedisApiInstance) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

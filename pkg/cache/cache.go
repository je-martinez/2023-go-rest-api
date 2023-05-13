package cache

import (
	"context"
	"time"

	"github.com/je-martinez/2023-go-rest-api/config"

	"github.com/redis/go-redis/v9"
)

func New(ctx *context.Context, cfg *config.RedisConfig) *RedisApiInstance {
	return &RedisApiInstance{
		ctx: ctx,
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
	ctx    *context.Context
	client *redis.Client
}

func (r *RedisApiInstance) Add(key string, value interface{}) error {
	return r.client.Set(*r.ctx, key, value, 0).Err()
}

func (r *RedisApiInstance) Remove(key ...string) (int64, error) {
	return r.client.Del(*r.ctx, key...).Result()
}

func (r *RedisApiInstance) Get(key string) (string, error) {
	return r.client.Get(*r.ctx, key).Result()
}

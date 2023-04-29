package cache

import (
	"context"
	"time"

	"github.com/je-martinez/2023-go-rest-api/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedisClient(cfg *config.Config) {

	if RedisClient != nil {
		return
	}

	redisHost := cfg.Redis.RedisAddr

	if redisHost == "" {
		redisHost = ":6379"
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password, // no password set
		DB:           cfg.Redis.DB,       // use default DB
	})
}

func Add(key string, value interface{}) error {
	err := RedisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	return val, err
}

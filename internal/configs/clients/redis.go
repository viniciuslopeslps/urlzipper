package clients

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
	"urlzipper/internal/configs/clients/mocks"
	"urlzipper/internal/configs/env"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}

func NewRedisClient(config *env.RedisConfig) RedisClient {
	if config.UseMock {
		return mocks.NewRedisMock()
	}

	return redis.NewClient(&redis.Options{
		Addr: config.Address,
	})
}

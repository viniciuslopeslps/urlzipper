package mocks

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisMock struct {
	DB map[string]any
}

func NewRedisMock() *RedisMock {
	return &RedisMock{
		DB: make(map[string]any),
	}
}

func (r RedisMock) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	r.DB[key] = value
	return redis.NewStatusCmd(ctx, key, value, expiration)
}

func (r RedisMock) Get(ctx context.Context, key string) *redis.StringCmd {
	get := r.DB[key]
	marshal, _ := json.Marshal(get)
	cmd := redis.NewStringCmd(ctx, get)
	cmd.SetVal(string(marshal))
	return cmd
}

package env

import (
	"log/slog"
	"strings"
	"time"
)

type RedisConfig struct {
	Address string
	UseMock bool
	TTL     time.Duration
}

type Environment struct {
	RedisConfig RedisConfig
}

func GetEnvConfig(scope string) *Environment {
	if strings.ToLower(scope) == "prod" {
		slog.Info("Using prod environment")
		return GetProdEnv()
	}

	slog.Info("Using local environment")
	return GetLocalEnv()
}

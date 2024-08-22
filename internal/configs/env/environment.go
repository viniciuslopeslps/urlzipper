package env

import (
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
		return GetProdEnv()
	}

	return GetLocalEnv()
}

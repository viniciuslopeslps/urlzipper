package env

import "time"

func GetLocalEnv() *Environment {
	return &Environment{
		RedisConfig: RedisConfig{
			UseMock: true,
			TTL:     time.Hour * 24,
		},
	}
}

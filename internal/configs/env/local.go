package env

import "time"

func GetLocalEnv() *Environment {
	return &Environment{
		RedisConfig: RedisConfig{
			UseMock: false,
			TTL:     time.Hour * 24,
			Address: "localhost:6379",
		},
	}
}

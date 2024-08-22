package env

import "time"

func GetProdEnv() *Environment {
	return &Environment{
		RedisConfig: RedisConfig{
			Address: "todo",
			TTL:     time.Hour * 24,
		},
	}
}

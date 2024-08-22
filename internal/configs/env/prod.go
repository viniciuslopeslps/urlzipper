package env

import "time"

func GetProdEnv() *Environment {
	return &Environment{
		RedisConfig: RedisConfig{
			Address: "urlzipper-prod-qo4kvi.serverless.use1.cache.amazonaws.com:6379",
			TTL:     time.Hour * 24,
		},
	}
}

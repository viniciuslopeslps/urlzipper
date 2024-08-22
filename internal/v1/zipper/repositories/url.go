package repositories

import (
	"context"
	"urlzipper/internal/configs/clients"
	"urlzipper/internal/configs/env"
	"urlzipper/internal/v1/zipper/errors"
	"urlzipper/internal/v1/zipper/models/entities"
)

type URLRepository interface {
	Save(url *entities.URL) errors.ApiError
}

type urlRepository struct {
	redisClient clients.RedisClient
	config      *env.RedisConfig
}

func NewURLRepository(config *env.RedisConfig, redisClient clients.RedisClient) URLRepository {
	return &urlRepository{
		config:      config,
		redisClient: redisClient,
	}
}

func (repo *urlRepository) Save(url *entities.URL) errors.ApiError {
	err := repo.redisClient.Set(context.Background(), url.Hash, url, repo.config.TTL).Err()
	if err != nil {
		return errors.DatabaseCommunicationError
	}

	return nil
}

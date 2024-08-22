package repositories

import (
	"context"
	"encoding/json"
	"urlzipper/internal/configs/clients"
	"urlzipper/internal/configs/env"
	"urlzipper/internal/v1/zipper/errors"
	"urlzipper/internal/v1/zipper/models/entities"
)

type URLRepository interface {
	Save(url *entities.URL) errors.ApiError
	FindURL(hash *string) (*entities.URL, errors.ApiError)
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

func (repo *urlRepository) FindURL(hash *string) (*entities.URL, errors.ApiError) {
	resString, err := repo.redisClient.Get(context.Background(), *hash).Result()
	if err != nil {
		return nil, errors.DatabaseCommunicationError
	}

	var res entities.URL
	err = json.Unmarshal([]byte(resString), &res)
	if err != nil {
		return nil, errors.DatabaseCommunicationError
	}

	return &res, nil

}

package mappers

import (
	"strconv"
	"time"
	"urlzipper/internal/configs/env"
	"urlzipper/internal/v1/zipper/models/dto"
	"urlzipper/internal/v1/zipper/models/entities"
)

type URLMapper interface {
	MapToURL(hash uint64, url string) *entities.URL
	MapToURLResponse(url *entities.URL) *dto.URLResponse
}

type urlMapper struct {
	config *env.RedisConfig
}

func NewURLMapper(config *env.RedisConfig) URLMapper {
	return &urlMapper{
		config: config,
	}
}

func (m *urlMapper) MapToURL(hash uint64, url string) *entities.URL {
	return &entities.URL{
		URL:       url,
		Hash:      strconv.Itoa(int(hash)),
		CreatedAt: time.Now(),
		TTL:       m.config.TTL,
	}
}

func (m *urlMapper) MapToURLResponse(url *entities.URL) *dto.URLResponse {
	return &dto.URLResponse{
		URL:        url.URL,
		Hash:       url.Hash,
		CreatedAt:  url.CreatedAt,
		Expiration: m.getMinutesToExpire(url.CreatedAt, url.TTL),
	}
}

func (*urlMapper) getMinutesToExpire(createdAt time.Time, ttl time.Duration) *time.Time {
	expired := createdAt.Add(ttl)
	return &expired
}

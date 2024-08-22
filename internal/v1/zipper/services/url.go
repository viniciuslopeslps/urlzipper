package services

import (
	"github.com/alextanhongpin/base62"
	"urlzipper/internal/v1/zipper/errors"
	"urlzipper/internal/v1/zipper/mappers"
	"urlzipper/internal/v1/zipper/models/dto"
	"urlzipper/internal/v1/zipper/repositories"
)

type URLService interface {
	Compress(req *dto.URLRequest) (*dto.URLResponse, errors.ApiError)
}

type urlService struct {
	mapper mappers.URLMapper
	repo   repositories.URLRepository
}

func NewURLService(mapper mappers.URLMapper, repo repositories.URLRepository) URLService {
	return &urlService{
		mapper: mapper,
		repo:   repo,
	}
}

func (service *urlService) Compress(req *dto.URLRequest) (*dto.URLResponse, errors.ApiError) {
	hash := base62.Decode(req.URL)

	url := service.mapper.MapToURL(hash, req.URL)
	err := service.repo.Save(url)
	if err != nil {
		return nil, err
	}

	return service.mapper.MapToURLResponse(url), nil
}

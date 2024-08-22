package services

import (
	"github.com/alextanhongpin/base62"
	"strconv"
	"urlzipper/internal/v1/zipper/errors"
	"urlzipper/internal/v1/zipper/mappers"
	"urlzipper/internal/v1/zipper/models/dto"
	"urlzipper/internal/v1/zipper/repositories"
)

type URLService interface {
	Compress(req *dto.URLRequest) (*dto.URLResponse, errors.ApiError)
	FindURL(hash string) (*dto.URLResponse, errors.ApiError)
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

	existent, err := service.repo.FindURL(strconv.FormatUint(hash, 10))
	if err != nil {
		return nil, err
	}
	if existent != nil {
		return nil, errors.URLAlreadyExists
	}

	url := service.mapper.MapToURL(hash, req.URL)
	err = service.repo.Save(url)
	if err != nil {
		return nil, err
	}

	return service.mapper.MapToURLResponse(url), nil
}

func (service *urlService) FindURL(hash string) (*dto.URLResponse, errors.ApiError) {
	urlResponse, err := service.repo.FindURL(hash)
	if err != nil {
		return nil, err
	}

	return service.mapper.MapToURLResponse(urlResponse), nil
}

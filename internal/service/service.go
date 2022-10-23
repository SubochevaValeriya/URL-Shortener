package service

import (
	"URLShortener/internal/repository"
)

type URLs interface {
	AddURL(originalUrl string) (string, error)
	GetURL(shortURL string) (string, error)
}

type Service struct {
	URLs
}

func NewService(repos *repository.Repository) *Service {
	return &Service{newApiService(repos.URLs)}
}

package service

import (
	_ "github.com/SubochevaValeriya/URL-Shortener"
	"github.com/SubochevaValeriya/URL-Shortener/internal/repository"
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

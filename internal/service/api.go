package service

import (
	urls "URLShortener"
	"URLShortener/internal/repository"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type ApiService struct {
	repo repository.URLs
}

func newApiService(repo repository.URLs) *ApiService {
	return &ApiService{repo: repo}
}

func (s *ApiService) AddURL(originalURL string) (string, error) {
	urlT, err := validateURL(originalURL)
	if err != nil {
		return "", err
	}

	shortURL := randomShortURL()

	urlInfo := urls.UrlInfo{
		OriginalURL: urlT.String(),
		ShortURL:    shortURL,
		CreatedAt:   time.Now(),
		Visits:      0,
	}

	_, err = s.repo.AddURL(&urlInfo)

	fmt.Println(shortURL)
	return shortURL, err
}

func validateURL(originalURL string) (*url.URL, error) {
	urlT, err := url.Parse(originalURL)
	if err != nil {
		return urlT, fmt.Errorf("validation error: %w", err)
	}
	return urlT, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var maxLengthOfShortURL = 10

func randomShortURL() string {
	rand.Seed(time.Now().UnixNano())
	lenght := rand.Intn(maxLengthOfShortURL-3) + 3
	shortURL := make([]byte, lenght)
	for i := range shortURL {
		shortURL[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(shortURL)
}

func (s *ApiService) GetURL(shortURL string) (string, error) {
	URLInfo := urls.UrlInfo{}
	err := s.repo.GetURL(&URLInfo, shortURL)
	s.repo.IncreaseVisits(shortURL, URLInfo.Visits+1)
	fmt.Println(URLInfo)
	fmt.Println(err)
	return URLInfo.OriginalURL, err
}

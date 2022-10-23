package repository

import (
	urls "URLShortener"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLs interface {
	AddURL(urlInfo *urls.UrlInfo) (*mongo.InsertOneResult, error)
	GetURL(c *urls.UrlInfo, shortURL string) error
	IncreaseVisits(shortURL string, visits int)
}

type Repository struct {
	URLs
}

func NewRepository(db *MongoHandler, tables DbTables) *Repository {
	return &Repository{NewApiMongo(db, tables)}
}

package repository

import (
	urls "github.com/SubochevaValeriya/URL-Shortener"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLs interface {
	AddURL(urlInfo *urls.UrlInfo) (*mongo.InsertOneResult, error)
	GetURL(c *urls.UrlInfo, shortURL string) error
	IncreaseVisits(shortURL string, visits int) error
	CreateIndex(field string) error
	DeleteURL(shortURL string)
}

type Repository struct {
	URLs
}

func NewRepository(db *MongoHandler, tables DbTables) *Repository {
	return &Repository{NewApiMongo(db, tables)}
}

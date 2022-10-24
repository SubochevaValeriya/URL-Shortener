package repository

import (
	"context"
	"fmt"
	urls "github.com/SubochevaValeriya/URL-Shortener"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ApiMongo struct {
	db       *MongoHandler
	dbTables DbTables
}

type DbTables struct {
	CollectionName string
}

func NewApiMongo(db *MongoHandler, dbTables DbTables) *ApiMongo {
	return &ApiMongo{db: db,
		dbTables: dbTables}
}

// AddURL adds new record to URLs table
func (r *ApiMongo) AddURL(urlInfo *urls.UrlInfo) (*mongo.InsertOneResult, error) {
	collection := r.db.client.Database(r.db.Database).Collection(r.dbTables.CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, urlInfo)
	fmt.Println(urlInfo)
	return result, err
}

// GetURL loads record on filter "short URL"
func (r *ApiMongo) GetURL(c *urls.UrlInfo, shortURL string) error {
	collection := r.db.client.Database(r.db.Database).Collection(r.dbTables.CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"shorturl": shortURL}).Decode(c)
	return err
}

// IncreaseVisits changes quantity of visits
func (r *ApiMongo) IncreaseVisits(shortURL string, visits int) {
	collection := r.db.client.Database(r.db.Database).Collection(r.dbTables.CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection.UpdateOne(ctx, bson.M{"shorturl": shortURL}, bson.D{{"$set", bson.D{{"visits", visits}}}})
}

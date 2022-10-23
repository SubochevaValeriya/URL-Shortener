package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoHandler struct {
	client   *mongo.Client
	Database string
}

type MongoConfig struct {
	Host            string
	Port            string
	DefaultDatabase string
}

// mongosh
// MongoHandler Constructor
func NewHandler(mongoConfig MongoConfig) (*MongoHandler, error) {
	address := "mongodb://" + mongoConfig.Host + ":" + mongoConfig.Port
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		return nil, err
	}
	mongoDB := &MongoHandler{
		client:   cl,
		Database: mongoConfig.DefaultDatabase,
	}
	return mongoDB, nil
}

//func NewHandler(address string, defaultDatabase string) (*MongoHandler, error) {
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
//	if err != nil {
//		return nil, err
//	}
//	mongoDB := &MongoHandler{
//		client:   cl,
//		database: defaultDatabase,
//	}
//	return mongoDB, nil
//}

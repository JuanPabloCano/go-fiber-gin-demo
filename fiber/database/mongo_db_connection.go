package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() *mongo.Client {
	mongoUri := "mongodb://localhost:27017/fiber-api-demo"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))

	if err != nil {
		panic(err)
	}

	return client
}

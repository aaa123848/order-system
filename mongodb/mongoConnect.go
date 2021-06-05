package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, url string) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

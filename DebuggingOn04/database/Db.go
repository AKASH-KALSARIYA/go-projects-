package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Dbinstantance() *mongo.Client {
	uri := "mongodb://localhost:27017"
	clientOpts := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = Dbinstantance()

func OpenCollection(client *mongo.Client, Collection string) *mongo.Collection {
	collection := client.Database("restaurant").Collection(Collection)
	return collection

}

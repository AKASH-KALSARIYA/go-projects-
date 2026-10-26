package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("auth").Collection("users")

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("DB not connected")
	}

	log.Println("MongoDB Connected")
}

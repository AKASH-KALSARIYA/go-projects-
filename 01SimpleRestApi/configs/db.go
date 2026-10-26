package Configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connection of mongo
const connectionstring = "mongodb://localhost:27017"
const dbname = "Tasks"
const colName = "Taskschedual"

var Collection *mongo.Collection

// connection
func ConnectDB() {

	clientOptions := options.Client().ApplyURI(connectionstring)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MONGO CONNECTION SUCCESSFULLY!!")

	Collection = client.Database(dbname).Collection(colName)

	fmt.Println("Collection instance is ready!!")
}

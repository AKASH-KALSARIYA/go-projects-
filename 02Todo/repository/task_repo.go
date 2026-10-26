package repository

import (
	"context"
	"fmt"
	"log"
	"todo/connfigs"
	"todo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insterttask(task model.Task) {
	if connfigs.Collection == nil {
		log.Fatal("doesnt connect a database!!")
	}

	insert, err := connfigs.Collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Id: ", insert.InsertedID)
}

func Getalltask() []model.Task {
	var tasks []model.Task
	cursor, err := connfigs.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task model.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return tasks
}

func Deletealltask() int64 {

	result, err := connfigs.Collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}

func UpdateOneTask(id string, task model.Task) int64 {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Invalid ID")
	}

	filter := bson.M{"_id": objID}

	update := bson.M{
		"$set": bson.M{
			"name":   task.Name,
			"status": task.Status,
		},
	}

	result, err := connfigs.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return result.ModifiedCount
}

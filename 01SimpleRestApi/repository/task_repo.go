package repository

import (
	"context"
	"fmt"
	"log"

	configs "github.com/AKASH-KALSARIYA/mygolang/FirstProject/configs"
	"github.com/AKASH-KALSARIYA/mygolang/FirstProject/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all tasks
func Getalltask() []bson.M {

	if configs.Collection == nil {
		log.Fatal("Database not initialized")
	}

	get, err := configs.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var tasks []bson.M

	for get.Next(context.Background()) {
		var task bson.M
		err := get.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	defer get.Close(context.Background())

	return tasks
}

// Insert task
func Insterttask(tasks model.Task) {

	if configs.Collection == nil {
		log.Fatal("Database not initialized")
	}

	insert, err := configs.Collection.InsertOne(context.Background(), tasks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted ID:", insert.InsertedID)
}

// Update task
func Updatetask(tasksId string, task model.Task) model.Task {

	if configs.Collection == nil {
		log.Fatal("Database not initialized")
	}

	id, err := primitive.ObjectIDFromHex(tasksId)
	if err != nil {
		fmt.Println("INVALID OBJECTID:", err)
		return model.Task{}
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"name":   task.Name,
		"status": task.Status,
	}}

	result, err := configs.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Matched:", result.MatchedCount)
	fmt.Println("Modified:", result.ModifiedCount)

	var updatedTask model.Task
	configs.Collection.FindOne(context.Background(), filter).Decode(&updatedTask)

	return updatedTask
}

// Delete task
func Deletetask(taskId string) {

	if configs.Collection == nil {
		log.Fatal("Database not initialized")
	}

	id, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	delete, err := configs.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete count:", delete.DeletedCount)
}

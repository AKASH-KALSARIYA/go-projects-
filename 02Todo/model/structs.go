package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name   string             `json:"name"`
	Status bool               `json:"status"`
}

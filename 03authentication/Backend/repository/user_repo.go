package repository

import (
	"auth/config"
	"auth/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertUser(user model.User) error {
	_, err := config.Collection.InsertOne(context.Background(), user)
	return err
}

func FindUserByEmail(email string) (model.User, error) {
	var user model.User
	err := config.Collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return user, err
}

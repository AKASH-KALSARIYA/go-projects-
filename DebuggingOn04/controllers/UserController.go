package controllers

import (
	"context"
	"net/http"
	"time"
	"math"

	"restaurant-management-system/database"
	"restaurant-management-system/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		cursor, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
			return
		}
		var users []bson.M
		if err := cursor.All(ctx, &users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		userId := c.Param("user_id")
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashed := HAshPassWord(*user.Password)
		user.Password = &hashed
		user.Created_at = time.Now()
		user.Updated_at = time.Now()
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		if _, err := userCollection.InsertOne(ctx, user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_id": user.User_id})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		type LoginReq struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var req LoginReq
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var user models.User
		if err := userCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		ok, _ := VerifyPassWord(*user.Password, req.Password)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_id": user.User_id})
	}
}

func round(num float64) int {
	if num < 0 {
		return int(num - 0.5)
	}
	return int(num + 0.5)
}

func toFixed(num float64, pricision int) float64 {
	output := math.Pow(10, float64(pricision))
	return float64(round(num*output)) / output
}

func HAshPassWord(Password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(Password), 14)
	return string(bytes)
}

func VerifyPassWord(UserPassword string, ProvidePassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(UserPassword), []byte(ProvidePassword))
	if err != nil {
		return false, "password mismatch"
	}
	return true, ""
}

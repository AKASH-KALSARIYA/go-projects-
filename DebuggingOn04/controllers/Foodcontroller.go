package controllers

import (
	"context"
	"fmt"
	"net/http"
	"restaurant-management-system/database"
	"restaurant-management-system/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		RecordPerPage, err := strconv.Atoi(c.Query("RecordPerPage"))
		if err != nil || RecordPerPage < 1 {
			RecordPerPage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * RecordPerPage
		if q := c.Query("startIndex"); q != "" {
			if si, e := strconv.Atoi(q); e == nil {
				startIndex = si
			}
		}

		opts := options.Find().SetLimit(int64(RecordPerPage)).SetSkip(int64(startIndex))
		cursor, err := foodCollection.Find(ctx, bson.M{}, opts)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listing foods"})
			return
		}
		var foods []bson.M
		if err := cursor.All(ctx, &foods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading foods"})
			return
		}
		c.JSON(http.StatusOK, foods)
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": " Error Occured While Feching The  Food Item!! "})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		if food.Menu_id == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "menu_id is required"})
			return
		}
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": *food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("menu is not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		food.Created_at = time.Now()
		food.Updated_at = time.Now()
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			msg := fmt.Sprintf("Food Item was not created!!")
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		foodId := c.Param("Food_id")
		var food models.Food
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj bson.D
		if food.Name != nil {
			updateObj = append(updateObj, bson.E{Key: "name", Value: *food.Name})
		}
		if food.Price != nil {
			num := toFixed(*food.Price, 2)
			updateObj = append(updateObj, bson.E{Key: "price", Value: num})
		}
		if food.Food_image != nil {
			updateObj = append(updateObj, bson.E{Key: "food_image", Value: *food.Food_image})
		}
		food.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: food.Updated_at})
		upsert := true
		opts := options.UpdateOptions{Upsert: &upsert}
		res, err := foodCollection.UpdateOne(ctx, bson.M{"food_id": foodId}, bson.D{{Key: "$set", Value: updateObj}}, &opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Food update failed"})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		foodId := c.Param("Food_id")
		_, err := foodCollection.DeleteOne(ctx, bson.M{"food_id": foodId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Food delete failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"deleted": foodId})
	}
}

func DeleteALLFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		_, err := foodCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Food delete all failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

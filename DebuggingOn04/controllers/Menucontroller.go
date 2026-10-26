package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant-management-system/database"
	"restaurant-management-system/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func Getmenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error Occured While Listing the menu!!"})
		}
		var allmenus []bson.M
		if err = result.All(ctx, &allmenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allmenus)
	}
}

func Getmenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("menu_id")
		var menu models.Menu
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": " Error Occured While Feching The  Food Item!! "})
		}
		c.JSON(http.StatusOK, menu)
	}
}

func Createmenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, inserterr := menuCollection.InsertOne(ctx, menu)
		if inserterr != nil {
			msg := fmt.Sprintf("menu item is not created!!")
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
		defer cancel()

	}
}

func intimespam(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}

func Updatemenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}
		var updateobj primitive.D
		if menu.Start_date != nil && menu.End_date != nil {
			if !intimespam(*menu.Start_date, *menu.End_date, time.Now()) {
				msg := "kindly retype a time"
				c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
				defer cancel()
				return
			}
			updateobj = append(updateobj, bson.E{Key: "start_date", Value: menu.Start_date})
			updateobj = append(updateobj, bson.E{Key: "end_date", Value: menu.End_date})
		}

		if menu.Name != "" {
			updateobj = append(updateobj, bson.E{Key: "name", Value: menu.Name})
		}
		if menu.Category != "" {
			updateobj = append(updateobj, bson.E{Key: "category", Value: menu.Category})
		}
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateobj = append(updateobj, bson.E{Key: "updated_at", Value: menu.Updated_at})
		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := menuCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{Key: "$set", Value: updateobj}},
			&opt,
		)
		if err != nil {
			msg := "Menu Update Failed!!"
			c.JSON(http.StatusInternalServerError, gin.H{"Error": msg})
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func Deletemenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		menuId := c.Param("menu_id")
		_, err := menuCollection.DeleteOne(ctx, bson.M{"menu_id": menuId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Menu delete failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"deleted": menuId})
	}
}

func DeleteALLmenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		_, err := menuCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Menu delete all failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}


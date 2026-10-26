package controllers

import (
	"context"
	"net/http"
	"time"

	"restaurant-management-system/database"
	"restaurant-management-system/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		cursor, err := tableCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list tables"})
			return
		}
		var tables []bson.M
		if err := cursor.All(ctx, &tables); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read tables"})
			return
		}
		c.JSON(http.StatusOK, tables)
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		tableId := c.Param("Order_id")
		var table models.Table
		err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "table not found"})
			return
		}
		c.JSON(http.StatusOK, table)
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var table models.Table
		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validate.Struct(table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		table.ID = primitive.NewObjectID()
		table.Table_id = table.ID.Hex()
		table.Created_at = time.Now()
		table.Updated_at = time.Now()
		if _, err := tableCollection.InsertOne(ctx, table); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create table"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"table_id": table.Table_id})
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		tableId := c.Param("Order_id")
		var table models.Table
		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj bson.D
		if table.Number_of_guest != nil {
			updateObj = append(updateObj, bson.E{Key: "number_of_guest", Value: *table.Number_of_guest})
		}
		if table.Table_number != nil {
			updateObj = append(updateObj, bson.E{Key: "table_number", Value: *table.Table_number})
		}
		table.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: table.Updated_at})
		upsert := true
		opts := options.UpdateOptions{Upsert: &upsert}
		res, err := tableCollection.UpdateOne(ctx, bson.M{"table_id": tableId}, bson.D{{Key: "$set", Value: updateObj}}, &opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update table"})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func DeleteTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		tableId := c.Param("Order_id")
		_, err := tableCollection.DeleteOne(ctx, bson.M{"table_id": tableId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete table"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"deleted": tableId})
	}
}

func DeleteALLTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		_, err := tableCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete tables"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

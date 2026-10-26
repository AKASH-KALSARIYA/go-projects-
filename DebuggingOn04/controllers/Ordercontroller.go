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

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		cursor, err := orderCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list orders"})
			return
		}
		var orders []bson.M
		if err := cursor.All(ctx, &orders); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read orders"})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		orderId := c.Param("Order_id")
		var order models.Order
		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var order models.Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validate.Struct(order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()
		order.Created_at = time.Now()
		order.Updated_at = time.Now()
		if _, err := orderCollection.InsertOne(ctx, order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"order_id": order.Order_id})
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		orderId := c.Param("Order_id")
		var order models.Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj bson.D
		if !order.Order_date.IsZero() {
			updateObj = append(updateObj, bson.E{Key: "order_date", Value: order.Order_date})
		}
		if order.Table_id != nil {
			updateObj = append(updateObj, bson.E{Key: "table_id", Value: *order.Table_id})
		}
		order.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: order.Updated_at})
		upsert := true
		opts := options.UpdateOptions{Upsert: &upsert}
		res, err := orderCollection.UpdateOne(ctx, bson.M{"order_id": orderId}, bson.D{{Key: "$set", Value: updateObj}}, &opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order"})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		orderId := c.Param("Order_id")
		_, err := orderCollection.DeleteOne(ctx, bson.M{"order_id": orderId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete order"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"deleted": orderId})
	}
}

func DeleteALLOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		_, err := orderCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete orders"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

package controllers

import (
	"context"
	"net/http"
	"time"

	"restaurant-management-system/database"
	"restaurant-management-system/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "order_item")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		cursor, err := orderItemCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list order items"})
			return
		}
		var items []bson.M
		if err := cursor.All(ctx, &items); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read order items"})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		id := c.Param("OrderItem_id")
		var item models.OrderItem
		err := orderItemCollection.FindOne(ctx, bson.M{"order_item_id": id}).Decode(&item)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "order item not found"})
			return
		}
		c.JSON(http.StatusOK, item)
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("OrderItem_id")
		items, err := IemByOrder(orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch order items"})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func IemByOrder(id string) (OrderItem []primitive.M, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	cursor, err := orderItemCollection.Find(ctx, bson.M{"order_id": id})
	if err != nil {
		return nil, err
	}
	var items []primitive.M
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var item models.OrderItem
		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validate.Struct(item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item.ID = primitive.NewObjectID()
		item.Order_item_id = item.ID.Hex()
		item.Created_at = time.Now()
		item.Updated_at = time.Now()
		if _, err := orderItemCollection.InsertOne(ctx, item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order item"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"order_item_id": item.Order_item_id})
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		id := c.Param("OrderItem_id")
		var item models.OrderItem
		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj bson.D
		if item.Quantity != nil {
			updateObj = append(updateObj, bson.E{Key: "quantity", Value: *item.Quantity})
		}
		if item.Unit_price != nil {
			updateObj = append(updateObj, bson.E{Key: "unit_price", Value: *item.Unit_price})
		}
		if item.Food_id != nil {
			updateObj = append(updateObj, bson.E{Key: "food_id", Value: *item.Food_id})
		}
		if item.Order_id != "" {
			updateObj = append(updateObj, bson.E{Key: "order_id", Value: item.Order_id})
		}
		item.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: item.Updated_at})
		upsert := true
		opts := options.UpdateOptions{Upsert: &upsert}
		res, err := orderItemCollection.UpdateOne(ctx, bson.M{"order_item_id": id}, bson.D{{Key: "$set", Value: updateObj}}, &opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order item"})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func DeleteOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		id := c.Param("OrderItem_id")
		_, err := orderItemCollection.DeleteOne(ctx, bson.M{"order_item_id": id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete order item"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"deleted": id})
	}
}

func DeleteALLOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		_, err := orderItemCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete order items"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

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

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func Getinvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		cursor, err := invoiceCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list invoices"})
			return
		}
		var invoices []bson.M
		if err := cursor.All(ctx, &invoices); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read invoices"})
			return
		}
		c.JSON(http.StatusOK, invoices)
	}
}

func Getinvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		id := c.Param("invoice_id")
		var invoice models.Invoice
		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": id}).Decode(&invoice)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
			return
		}
		c.JSON(http.StatusOK, invoice)
	}
}

func Createinvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var invoice models.Invoice
		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validate.Struct(invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		invoice.ID = primitive.NewObjectID()
		invoice.Invoice_id = invoice.ID.Hex()
		invoice.Created_at = time.Now()
		invoice.Updated_at = time.Now()
		if _, err := invoiceCollection.InsertOne(ctx, invoice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create invoice"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"invoice_id": invoice.Invoice_id})
	}
}

func Updateinvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		id := c.Param("invoice_id")
		var invoice models.Invoice
		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj bson.D
		if invoice.Order_id != "" {
			updateObj = append(updateObj, bson.E{Key: "order_id", Value: invoice.Order_id})
		}
		if invoice.Payment_Method != nil {
			updateObj = append(updateObj, bson.E{Key: "payment_method", Value: *invoice.Payment_Method})
		}
		if invoice.Payment_Status != nil {
			updateObj = append(updateObj, bson.E{Key: "payment_status", Value: *invoice.Payment_Status})
		}
		if !invoice.Payment_Due_Date.IsZero() {
			updateObj = append(updateObj, bson.E{Key: "payment_due_date", Value: invoice.Payment_Due_Date})
		}
		invoice.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: invoice.Updated_at})
		upsert := true
		opts := options.UpdateOptions{Upsert: &upsert}
		res, err := invoiceCollection.UpdateOne(ctx, bson.M{"invoice_id": id}, bson.D{{Key: "$set", Value: updateObj}}, &opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update invoice"})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func Deleteinvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		id := c.Param("invoice_id")
		_, err := invoiceCollection.DeleteOne(ctx, bson.M{"invoice_id": id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete invoice"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"deleted": id})
	}
}

func DeleteALLinvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		_, err := invoiceCollection.DeleteMany(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete invoices"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/user"
	"time"

	"github.com/gin-gonic/gin"
)

	type Application struct {
		prodCollection *mongo.Collection
		userCollection *mongo.Collection
	}

	func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
		return &Application{prodCollection: prodCollection,userCollection: userCollection}
	}
func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("Product ID is missing")
			_ = c.AbortWithStatusJSON(http.StatusBadRequest, errors.New{"error": "Product ID is missing"})
			return
		}

		useQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("User ID is missing")
			_ = c.AbortWithStatusJSON(http.StatusBadRequest, errors.New{"error": "User ID is missing"})
			return
		}

		primitive.ObjectIDFromHex(productQueryID)
}

func RemoteItem() gin.HandlerFunc {
}

func GetItemFromCart() gin.HandlerFunc {
}

func BuyFromCart() gin.HandlerFunc {
}

func InstantBuy() gin.HandlerFunc {
}


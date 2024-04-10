package controllers

import (
	"context"
	"errors"
	"go-ecommerce/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}

}
func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		prodQueryID := c.Query("id")
		if prodQueryID == "" {
			log.Println("Product ID empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product ID empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("User ID empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("User ID empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(prodQueryID)

		if err != nil {
			log.Panicln(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		err = database.AddToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Add to cart successfully")
	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		prodQueryID := c.Query("id")
		if prodQueryID == "" {
			log.Println("Product ID empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product ID empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("User ID empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("User ID empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(prodQueryID)

		if err != nil {
			log.Panicln(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		err = database.RemoveFromCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Removed Item Successfully")
	}
}

func GetItemInCart() gin.HandlerFunc {

}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")

		if userQueryID == "" {
			log.Panicln("user id empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("User ID empty"))

			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

			defer cancel()
		}

		err := database.BuyFromCart(ctx, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON("Order setup")
	}
}

func (app *Application) BuyNow() gin.HandlerFunc {
	return func(c *gin.Context) {
		prodQueryID := c.Query("id")
		if prodQueryID == "" {
			log.Println("Product ID empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product ID empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("User ID empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("User ID empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(prodQueryID)

		if err != nil {
			log.Panicln(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		err = databse.BuyNow(ctx.app.prodCollection, app.userCollection, productID, prodQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Order Setup")
	}
}

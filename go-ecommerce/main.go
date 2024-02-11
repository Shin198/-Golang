package main

import (
	"go-ecommerce/controllers"
	"go-ecommerce/database"
	"go-ecommerce/middleware"
	"go-ecommerce/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	app := controllers.newApp(database.ProductData(database.Client, "Product"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removefromcart", app.RemoveFromCart())
	router.GET("/checkoutcart", app.BuyFromCart())
	router.GET("/buynow", app.BuyNow())

	log.Fatal(router.Run(":" + port))
}

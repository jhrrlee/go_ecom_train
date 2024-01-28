package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jhrrlee/go_ecom_train/controllers"
	"github.com/jhrrlee/go_ecom_train/database"
	"github.com/jhrrlee/go_ecom_train/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "products"), database.UserData(database.Client, "users"))

	router := gin.New()
	router.Use(gin.Logger())

	router.UserRoute(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromeCart())
	router.GET("/instantbuy", app.instantbuy())

	log.Fatal(router.Run(":" + port))
}

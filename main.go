package main

import (
	"os"
	"github.com/gin-gonic/gin"
  "golang-restaurant-management/controllers"
  "golang-restaurant-management/database"
  "golang-restaurant-management/middleware"
  "golang-restaurant-management/routes"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
  routes.UserRoutes(router)
  router.Use(middleware.Authentication())

  routes.FoodRoutes(router)
  routes.MenuRoutes(router)
  routes.TableRoutes(router)
  routes.OrderRoutes(router)
  routes.OrderItemRoutes(router)
  routes.InvoiceRoutes(router)

  router.Run(":" + port)
}
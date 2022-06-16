package main

import (
	"os"
	"github.com/gin-gonic/gin"
  "golang-restaurant-management/controllers"
  "golang-restaurant-management/database"
  "golang-restaurant-management/middleware"
  "golang-restaurant-management/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

}
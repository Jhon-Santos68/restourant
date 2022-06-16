package controller 

import (
  "context"
  "golang-restaurant-management/database"
  "golang-restaurant-management/models"
  "net/http"
  "time"
  
  "github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/mongo"
  "gopkg.in/bluesuncorp/validator.v5"
  "go.mongodb.org/mongo-driver/bson"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc{
  return func(c *gin.Context){
    
  }
}

func GetFood() gin.HandlerFunc{
  return func(c *gin.Context){
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
    foodId := c.Param("food_id")
    var food models.Food

    err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
    defer cancel()
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
    }

    c.JSON(http.StatusOK, food)
  }
}

func CreateFood() gin.HandlerFunc{
  return func(c *gin.Context){
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
    var menu models.Menu
    var food models.Food

    if err := c.BindJSON(&food); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
    }

    validate.Struct
  }
}

func round(num float64) int {
  
}

func toFixed(num float64, precision int) float64 {
  
}

func UpdateFood() gin.HandlerFunc{
  return func(c *gin.Context){
    
  }
}

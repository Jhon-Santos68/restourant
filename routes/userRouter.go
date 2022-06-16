package routes

import (
  "github.com/gin-gonic/gin"
  controller "golang-restaurant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
  incomingRoutes.GET("/users", controller.GetUser())
  incomingRoutes.GET("/users/:user_id", controller.GetUser())
  incomingRoutes.POST("/users/singup", controller.SingUp())
  incomingRoutes.POST("/users/login", controller.Login())
  
}
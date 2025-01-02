package routes

import (
    "github.com/gin-gonic/gin"
    "crm-app/internal/controllers"
)

func SetupRoutes(router *gin.Engine) {
    userController := controllers.UserController{}

    router.POST("/users", userController.CreateUser)
    router.GET("/users/:id", userController.GetUser)
    router.PUT("/users/:id", userController.UpdateUser)
    router.DELETE("/users/:id", userController.DeleteUser)
}
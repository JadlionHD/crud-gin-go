package main

import (
	"github.com/JadlionHD/crud-gin-go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/posts/:id", controllers.GetPost)
	router.POST("/posts", controllers.CreatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
	router.PUT("/posts/:id", controllers.UpdatePost)

	router.Run("localhost:8080")
}

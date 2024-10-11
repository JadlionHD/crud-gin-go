package main

import (
	"fmt"
	"net/http"

	"github.com/JadlionHD/crud-gin-go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		data, err := controllers.GetPostDummies(id)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Posts %s not found", id),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	router.POST("/posts/:id", func(ctx *gin.Context) {

	})
	router.Run("localhost:8080")
}

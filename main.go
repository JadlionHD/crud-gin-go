package main

import (
	"fmt"
	"net/http"

	"github.com/JadlionHD/crud-gin-go/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		data, err := utils.GetDummies(id)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Posts %s not found", id),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	})
	router.Run("localhost:8080")
}

package routes

import (
	"fmt"
	"net/http"

	"github.com/JadlionHD/crud-gin-go/api/server/controllers"
	"github.com/JadlionHD/crud-gin-go/api/server/types"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	id := c.Param("id")
	data, err := controllers.GetDummies(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Posts %s not found", id),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func CreatePost(c *gin.Context) {

	var jsonData types.CreatePostInput

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error when parsing json",
		})
		return
	}

	err := controllers.CreateDummies(jsonData.ID, jsonData.Title, jsonData.Body, jsonData.Tags)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error when creating data",
		})
		return
	}

}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var jsonData types.CreatePostInput

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error when parsing json",
		})
		return
	}

	ok, err := controllers.UpdateDummies(id, jsonData.Title, jsonData.Body, jsonData.Tags)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("internal error: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("sucessfully update  post with id: %s", id),
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	ok := controllers.DeleteDummies(id)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail to delete posts",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("sucessfully delete post with id: %s", id),
	})
}

package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/JadlionHD/crud-gin-go/utils"
	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int64         `json:"id"`
	Title     string        `json:"title"`
	Body      string        `json:"body"`
	Tags      []string      `json:"tags"`
	Reactions PostReactions `json:"reactions"`
	Views     int           `json:"views"`
	UserID    int           `json:"userId"`
}

type PostReactions struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

type CreatePostInput struct {
	ID    int64    `json:"id" binding:"required"`
	Title string   `json:"title" binding:"required"`
	Body  string   `json:"body" binding:"required"`
	Tags  []string `json:"tags" binding:"required"`
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	data, err := GetDummies(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Posts %s not found", id),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func CreatePost(c *gin.Context) {

	var jsonData CreatePostInput

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error when parsing json",
		})
		return
	}

	err := CreateDummies(jsonData.ID, jsonData.Title, jsonData.Body, jsonData.Tags)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error when creating data",
		})
		return
	}

}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var jsonData CreatePostInput

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal error when parsing json",
		})
		return
	}

	ok, err := UpdateDummies(id, jsonData.Title, jsonData.Body, jsonData.Tags)

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
	ok := DeleteDummies(id)

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

func GetDummies(id string) (*Post, error) {
	result, err := utils.ReadJSON[Post](fmt.Sprintf("dummies/posts/%s.json", id))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateDummies(id string, title string, body string, tags []string) (bool, error) {
	data, err := GetDummies(id)
	parsedId, _ := strconv.ParseInt(id, 10, 32)

	if err != nil {
		return false, err
	}

	ok := DeleteDummies(id)

	if !ok {
		return false, errors.New("fail to update dummies")
	}

	newData := Post{
		ID:    parsedId,
		Title: title,
		Body:  body,
		Tags:  tags,
		Reactions: PostReactions{
			Likes:    data.Reactions.Likes,
			Dislikes: data.Reactions.Dislikes,
		},
		Views:  data.Views,
		UserID: data.UserID,
	}

	written := utils.WriteJSON(fmt.Sprintf("dummies/posts/%d.json", parsedId), newData)
	if !written {
		return false, errors.New("file not written")
	}

	return true, nil
}

func DeleteDummies(id string) bool {
	parsedId, _ := strconv.ParseInt(id, 10, 32)
	err := os.Remove(fmt.Sprintf("dummies/posts/%d.json", parsedId))

	return err == nil
}

func CreateDummies(id int64, title string, body string, tags []string) error {
	data := Post{
		ID:    id,
		Title: title,
		Body:  body,
		Tags:  tags,
		Reactions: PostReactions{
			Likes:    0,
			Dislikes: 0,
		},
		Views:  0,
		UserID: 0,
	}

	written := utils.WriteJSON(fmt.Sprintf("dummies/posts/%d.json", id), data)
	if !written {
		return errors.New("file not written")
	}

	return nil
}

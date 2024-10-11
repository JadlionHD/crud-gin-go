package controllers

import (
	"fmt"

	"github.com/JadlionHD/crud-gin-go/utils"
)

type Post struct {
	ID        int           `json:"id"`
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
}

func GetPostDummies(id string) (*Post, error) {
	result, err := utils.ReadJSON[Post](fmt.Sprintf("dummies/posts/%s.json", id))

	if err != nil {
		return nil, err
	}

	return result, nil
}

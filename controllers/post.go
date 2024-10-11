package controllers

import (
	"errors"
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

func GetPostDummies(id string) (*Post, error) {
	result, err := utils.ReadJSON[Post](fmt.Sprintf("dummies/posts/%s.json", id))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateDummies(id int, title string, body string, tags []string) error {
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

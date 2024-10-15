package controllers

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/JadlionHD/crud-gin-go/api/server/types"
	"github.com/JadlionHD/crud-gin-go/api/utils"
)

func GetDummies(id string) (*types.Post, error) {
	result, err := utils.ReadJSON[types.Post](fmt.Sprintf("test/dummies/posts/%s.json", id))

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

	newData := types.Post{
		ID:    parsedId,
		Title: title,
		Body:  body,
		Tags:  tags,
		Reactions: types.PostReactions{
			Likes:    data.Reactions.Likes,
			Dislikes: data.Reactions.Dislikes,
		},
		Views:  data.Views,
		UserID: data.UserID,
	}

	written := utils.WriteJSON(fmt.Sprintf("test/dummies/posts/%d.json", parsedId), newData)
	if !written {
		return false, errors.New("file not written")
	}

	return true, nil
}

func DeleteDummies(id string) bool {
	parsedId, _ := strconv.ParseInt(id, 10, 32)
	err := os.Remove(fmt.Sprintf("test/dummies/posts/%d.json", parsedId))

	return err == nil
}

func CreateDummies(id int64, title string, body string, tags []string) error {
	data := types.Post{
		ID:    id,
		Title: title,
		Body:  body,
		Tags:  tags,
		Reactions: types.PostReactions{
			Likes:    0,
			Dislikes: 0,
		},
		Views:  0,
		UserID: 0,
	}

	written := utils.WriteJSON(fmt.Sprintf("test/dummies/posts/%d.json", id), data)
	if !written {
		return errors.New("file not written")
	}

	return nil
}

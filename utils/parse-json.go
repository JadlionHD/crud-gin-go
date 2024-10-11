package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/JadlionHD/crud-gin-go/controllers"
)

func GetDummies(id string) (controllers.Post, error) {
	jsonFile, err := os.Open(fmt.Sprintf("dummies/posts/%s.json", id))

	if err != nil {
		fmt.Println(err)
		return controllers.Post{}, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result controllers.Post

	json.Unmarshal(byteValue, &result)

	return result, nil
}

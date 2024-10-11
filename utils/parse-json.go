package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/JadlionHD/crud-gin-go/controllers"
)

func GetDummies(id string) (*controllers.Post, error) {
	result, err := ReadJSON[controllers.Post](fmt.Sprintf("dummies/posts/%s.json", id))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func ReadJSON[T any](path string) (*T, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result *T

	json.Unmarshal(byteValue, &result)

	return result, nil
}

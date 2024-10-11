package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

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

func WriteJSON(filePath string, data any) bool {
	file, err := json.Marshal(data)

	os.WriteFile(filePath, file, 0644)

	return err == nil
}

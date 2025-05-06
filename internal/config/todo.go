package config

import (
	"encoding/json"
	"fmt"
	"go-todo/internal/logger"
	"go-todo/internal/models"
	"os"
)

type TodoList []models.TodoTask

/*
Checks to see if json config exists, if not it creates it
*/
func (t *TodoList) Init(filePath string) error {

	// stat file if not exist then
	var _, error = os.Stat(filePath)

	// Need to create File
	if error != nil {
		fmt.Println("No todo.json... Creating File.")
		var file, fileErr = os.Create(filePath)

		if fileErr != nil {
			logger.LogError(fileErr)
			return fileErr
		}

		var byte, jsonErr = json.Marshal(t)

		if jsonErr != nil {
			logger.LogError(jsonErr)
			return jsonErr
		}

		_, error = file.Write(byte)
		if error != nil {
			logger.LogError(error)
			return error
		}
		error = file.Sync()
		if error != nil {
			logger.LogError(error)
		}

		error = file.Close()
		if error != nil {
			logger.LogError(error)
		}

		return nil
	}

	var byte, readErr = os.ReadFile(filePath)

	if readErr != nil {
		logger.LogError(readErr)
		return readErr
	}

	json.Unmarshal(byte, &t)

	return nil
}

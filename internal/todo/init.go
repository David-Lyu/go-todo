package todo

import (
	"encoding/json"
	"fmt"
	"go-todo/internal/logger"
	"os"
)

/*
Checks to see if json config exists, if not it creates it
*/
func (t *Todo) Init(filePath string) error {

	// stat file if not exist then
	var _, error = os.Stat(filePath)

	// Need to create File
	if error != nil {
		fmt.Println("No todo.json... Creating File.")
		var file, fileErr = os.Create(filePath)

		if fileErr != nil {
			return logger.HandleError(fileErr, true)
		}

		var byte, jsonErr = json.Marshal(t)

		if jsonErr != nil {
			return logger.HandleError(jsonErr, true)
		}

		_, error = file.Write(byte)
		if error != nil {
			return logger.HandleError(error, true)
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
		return logger.HandleError(readErr, true)
	}

	json.Unmarshal(byte, &t.TodoList)

	return nil
}

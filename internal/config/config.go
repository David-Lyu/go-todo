package config

import (
	"encoding/json"
	"fmt"
	"go-todo/internal/logger"
	"os"
)

type Config struct {
	TodoPath  string
	IsCreated bool
}

func (c Config) Init() error {

	var configDirPath = ""
	var configFilePath = "config.json"
	var error error
	var file *os.File

	configDirPath, error = os.UserHomeDir()
	if error != nil {
		//handle error here
		logger.LogError("Error getting User Dir")
		return error
	}

	configDirPath += string(os.PathSeparator) + ".external_configs"
	configDirPath += string(os.PathSeparator) + "go_todo" + string(os.PathSeparator)

	// Check if ~/.external_configs/go_todo/config.json
	_, error = os.Stat(configDirPath + configFilePath)

	if error != nil {
		fmt.Println("Creating ~/.external_config/go_todo/config.json")

		var _, dirError = os.Stat(configDirPath)

		// Full dir doesnt exist, will make the directory
		if dirError != nil {
			var mkdirError = os.MkdirAll(configDirPath, 0750)
			if mkdirError != nil {
				logger.LogError(mkdirError)
				return error
			}
		}

		//Creating file if dir exists
		file, error = os.Create(configDirPath + configFilePath)
		c.TodoPath = configDirPath + "todo.json"
		c.IsCreated = true

		var configContents, jsonError = json.Marshal(c)

		if jsonError != nil {
			logger.LogError(jsonError)
			return jsonError
		}

		file.Write(configContents)

	}

	return nil
}

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

		if _, dirError := os.Stat(configDirPath); dirError != nil {
			// Full dir doesnt exist, will make the directory
			if funcErr := createDir(configDirPath); funcErr != nil {
				return funcErr
			}
		}
		// creating missing file
		file, error = initFile(configDirPath, configFilePath, &c)

		if error != nil {
			return error
		}
	}

	if file == nil {
		os.Open(configDirPath + configFilePath)
	}

	return nil
}

func createDir(configDirPath string) error {
	var mkdirError = os.MkdirAll(configDirPath, 0750)
	if mkdirError != nil {
		logger.LogError(mkdirError)
		return mkdirError
	}
	return nil
}

func initFile(dirPath string, filePath string, c *Config) (*os.File, error) {
	file, error := os.Create(dirPath + filePath)

	if error != nil {
		logger.LogError(error)
		return file, error
	}

	c.TodoPath = dirPath + "todo.json"
	c.IsCreated = true

	var configContents, jsonError = json.Marshal(c)

	if jsonError != nil {
		logger.LogError(jsonError)
		return file, jsonError
	}
	//Creating file if dir exists
	file.Write(configContents)

	return file, nil
}

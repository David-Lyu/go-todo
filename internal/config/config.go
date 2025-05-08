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

func (c *Config) Init() error {

	var configDirPath = ""
	var configFilePath = "config.json"
	var error error

	configDirPath, error = os.UserHomeDir()
	if error != nil {
		return logger.HandleError(error, true)
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
		error = initFile(configDirPath, configFilePath, c)

		if error != nil {
			return error
		}
	}
	// No error means file exists.
	if c.TodoPath == "" {
		buff, readErr := os.ReadFile(configDirPath + configFilePath)
		if readErr != nil {
			return logger.HandleError(error, true)
		}

		unmarshalErr := json.Unmarshal(buff, &c)
		if unmarshalErr != nil {
			return logger.HandleError(unmarshalErr, true)
		}

		if c.TodoPath == "" {
			panic("Todo Path needs to be set in config file")
		}
		c.IsCreated = true
	}

	return nil
}

// Is this even needed? Only if todopath is
func (c Config) GetTodoPath() string {
	return c.TodoPath
}

func createDir(configDirPath string) error {
	var mkdirError = os.MkdirAll(configDirPath, 0750)
	if mkdirError != nil {
		logger.LogError(mkdirError)
		return mkdirError
	}
	return nil
}

func initFile(dirPath string, filePath string, c *Config) error {
	file, error := os.Create(dirPath + filePath)

	if error != nil {
		logger.LogError(error)
		return error
	}

	c.TodoPath = dirPath + "todo.json"
	c.IsCreated = true

	var configContents, jsonError = json.Marshal(c)

	if jsonError != nil {
		logger.LogError(jsonError)
		return jsonError
	}
	//Creating file if dir exists
	_, error = file.Write(configContents)
	if error != nil {
		logger.LogError(error)
		return error
	}
	_, error = file.Seek(0, 0)
	if error != nil {
		logger.LogError(error)
		return error
	}
	error = file.Close()
	if error != nil {
		logger.LogError(error)
		return error
	}

	return nil
}

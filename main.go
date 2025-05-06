package main

import (
	"fmt"
	"go-todo/internal/config"
	"go-todo/internal/logger"
	"log"
)

func main() {
	var shouldExit = false

	// install configs
	var conf = config.Config{}
	var error = conf.Init()
	if error != nil {
		log.Fatal(error)
	}

	//grab todos
	var todo = config.TodoList{}
	error = todo.Init(conf.TodoPath)

	for shouldExit != true {
		var input string

		fmt.Println("Press 1 to add \nPress 0 to exit")
		var _, scanErr = fmt.Scanln(&input)
		if scanErr != nil {
			logger.LogError(scanErr)
			shouldExit = true
		}

		switch input {
		case "1":
			fmt.Println("You pressed 1")
		case "0":
			fmt.Println("time to exit")
			shouldExit = true
		default:
			fmt.Println("You pressed incorrect value")
		}
	}
}

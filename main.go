package main

import (
	"bufio"
	"fmt"
	"go-todo/internal/config"
	"go-todo/internal/logger"
	"go-todo/internal/todo"
	"log"
	"os"
)

func main() {
	var shouldExit = false

	// install configs
	var conf = config.Config{}
	var globalError = conf.Init()
	if globalError != nil {
		log.Fatal(globalError)
	}

	//grab todos
	var todoList = todo.TodoList{}
	globalError = todoList.Init(conf.TodoPath)
	fmt.Println(todoList)

	// var test []byte
	// var testError error
	// test, testError = io.ReadAll(os.Stdin)

	// if testError != nil {
	// 	fmt.Println(testError)
	// 	return
	// }

	// if test != nil {
	// 	return
	// }
	// fmt.Println(test)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		log.Println("line", s.Text())
	}

	for shouldExit != false {
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

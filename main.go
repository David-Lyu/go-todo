package main

import (
	"go-todo/internal/config"
	"go-todo/internal/menu"
	"go-todo/internal/todo"
	"log"
)

func main() {

	// install configs
	var conf = config.Config{}
	var globalError = conf.Init()
	if globalError != nil {
		log.Fatal(globalError)
	}

	//grab todos
	var todoList = todo.Todo{}
	globalError = todoList.Init(conf.TodoPath)

	var newMenu = menu.StartMenu{UserChoice: 1}
	newMenu.MenuStart(&todoList)

}

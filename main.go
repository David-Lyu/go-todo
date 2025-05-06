package main

import (
	"fmt"
	"go-todo/internal/config"
)

func main() {
	// install configs
	var config = config.Config{}
	config.Init()
	fmt.Println(config.TodoPath)
}

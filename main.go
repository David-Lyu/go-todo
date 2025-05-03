package main

import "go-todo/internal/config"

func main() {
	// install configs
	var config = config.Config{}
	config.Init()
}

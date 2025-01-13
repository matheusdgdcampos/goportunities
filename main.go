package main

import (
	"goportunities/config"
	"goportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
		return
	}
	router.Initialize()
}

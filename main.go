package main

import (
	"github.com/ricksantos88/goopportunities/config"
	"github.com/ricksantos88/goopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}

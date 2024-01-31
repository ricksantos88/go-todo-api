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

	// Initialize Router
	router.Initialize()
}

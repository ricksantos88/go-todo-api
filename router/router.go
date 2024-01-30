package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	router := gin.Default()
	var openingRouter = router.Group("/api/v1")

	// Initialize opening routes
	initializeOpeningRoutes(openingRouter)

	// Run the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

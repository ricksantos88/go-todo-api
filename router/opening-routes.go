package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ricksantos88/goopportunities/handler"
)

func initializeOpeningRoutes(router *gin.RouterGroup) {

	v1 := router
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}

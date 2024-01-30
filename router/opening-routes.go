package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ricksantos88/goopportunities/handler"
)

func initializeOpeningRoutes(router *gin.RouterGroup) {

	v1 := router
	{
		v1.GET("/opening", handler.ShowOpeningHendler)
		v1.POST("/opening", handler.CreateOpeningHendler)
		v1.PUT("/opening", handler.UpdateOpeningHendler)
		v1.DELETE("/opening", handler.DeleteOpeningHendler)
		v1.GET("/openings", handler.ListOpeningHendler)
	}

}

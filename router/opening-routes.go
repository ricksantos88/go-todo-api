package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ricksantos88/goopportunities/handler"
)

func initializeOpeningRoutes(router *gin.RouterGroup) {

	v1 := router
	{
		v1.GET("/todo", handler.ShowTodoHandler)
		v1.POST("/todo", handler.CreateTodoHandler)
		v1.PUT("/todo", handler.UpdateTodoHandler)
		v1.DELETE("/todo", handler.DeleteTodoHandler)
		v1.GET("/todos", handler.ListTodosHandler)
	}

}

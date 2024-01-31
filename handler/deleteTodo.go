package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTodoHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE opening",
	})
}

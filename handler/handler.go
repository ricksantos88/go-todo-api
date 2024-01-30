package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHendler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST opening",
	})
}

func ShowOpeningHendler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})
}

func UpdateOpeningHendler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT opening",
	})
}

func DeleteOpeningHendler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE opening",
	})
}

func ListOpeningHendler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET openings",
	})
}

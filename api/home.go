package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "EthioGrego Server",
	})
}

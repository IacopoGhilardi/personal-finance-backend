package userRoutes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {

	router.GET("/ping", func (context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
}


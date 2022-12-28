package userRoutes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {

	router.GET("/token", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

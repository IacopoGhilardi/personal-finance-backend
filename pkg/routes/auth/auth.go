package authRoutes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "auth",
		})
	})

	router.GET("/register", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "auth",
		})
	})

	router.GET("/login", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "auth",
		})
	})
}

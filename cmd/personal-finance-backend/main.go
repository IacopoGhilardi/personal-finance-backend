package main

import (
	// "os"
	"github.com/gin-gonic/gin"
	"github.com/IacopoGhilardi/personal-finance-backend/pkg/config"
	"github.com/IacopoGhilardi/personal-finance-backend/pkg/controllers"
	// "github.com/IacopoGhilardi/personal-finance-backend/pkg/models"
)

func main() {

	// Get Client, Context, CancelFunc and
    // err from connect method.
    client, ctx, cancel, err := config.Connect("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }

	// Release resource when the main
    // function is returned.
    defer config.Close(client, ctx, cancel)
     
    // Ping mongoDB with Ping method
    config.Ping(client, ctx)

	userController.Create(client, ctx)

	router := gin.Default()

	router.GET("/ping", func (context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
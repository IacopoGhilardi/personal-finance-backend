package main

import (
	"github.com/IacopoGhilardi/personal-finance-backend/pkg/config/database"
	"github.com/IacopoGhilardi/personal-finance-backend/pkg/controllers"
	"github.com/IacopoGhilardi/personal-finance-backend/pkg/routes"
)

func main() {

	// Get Client, Context, CancelFunc and
    // err from connect method.
    client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
    if err != nil {
        panic(err)
    }

	// Release resource when the main
    // function is returned.
    defer database.Close(client, ctx, cancel)
     
    // Ping mongoDB with Ping method
    database.Ping(client, ctx)

	userController.Create(client, ctx)

	indexRoutes.InitRoutes()
}
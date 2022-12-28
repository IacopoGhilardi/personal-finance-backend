package userController

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/IacopoGhilardi/personal-finance-backend/pkg/models"
)

func Create(client *mongo.Client, ctx context.Context) {
	fmt.Println("Create user")
	provaUser := models.User{
		Email: "prova@prova.com",
		Password: "Ciao",
	}

	database := client.Database("personal-finance")
	userCollection := database.Collection("users")

	insertResult, err := userCollection.InsertOne(ctx, provaUser) 
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)
}
package indexRoutes

import (
	authRoutes "github.com/IacopoGhilardi/personal-finance-backend/pkg/routes/auth"
	userRoutes "github.com/IacopoGhilardi/personal-finance-backend/pkg/routes/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {

	router := gin.Default()

	authGroup := router.Group("/auth")
	userGroup := router.Group("/user")

	authRoutes.InitRoutes(authGroup)
	userRoutes.InitRoutes(userGroup)

	router.Run()
}

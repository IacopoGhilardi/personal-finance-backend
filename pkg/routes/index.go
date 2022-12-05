package indexRoutes

import (
	"github.com/gin-gonic/gin"
	"github.com/IacopoGhilardi/personal-finance-backend/pkg/routes/user"
	"github.com/IacopoGhilardi/personal-finance-backend/pkg/routes/auth"
)

func InitRoutes() {

	router := gin.Default()

	userGroup := router.Group("/user")
	userRoutes.InitRoutes(userGroup)

	authGroup := router.Group("/auth")
	authRoutes.InitRoutes(authGroup)

	router.Run()
}
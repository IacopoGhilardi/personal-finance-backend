package indexRoutes

import (
	authRoutes "github.com/IacopoGhilardi/personal-finance-backend/pkg/routes/auth"
	plaidRoutes "github.com/IacopoGhilardi/personal-finance-backend/pkg/routes/plaid"
	userRoutes "github.com/IacopoGhilardi/personal-finance-backend/pkg/routes/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {

	router := gin.Default()

	authGroup := router.Group("/auth")
	userGroup := router.Group("/user")
	plaidGroup := router.Group("/plaid")

	authRoutes.InitRoutes(authGroup)
	userRoutes.InitRoutes(userGroup)
	plaidRoutes.InitRoutes(plaidGroup)

	router.Run()
}

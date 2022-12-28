package plaidRoutes

import (
	plaidController "github.com/IacopoGhilardi/personal-finance-backend/pkg/controllers/plaid"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {

	router.GET("/token", plaidController.GetAccessToken)
}

package userRoutes

import (
	bankController "github.com/IacopoGhilardi/personal-finance-backend/pkg/controllers/bank"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {

	router.GET("/token", bankController.GetAccessToken)
	router.GET("/bank/list", bankController.GetBankList)
}

package plaidController

import (
	"context"
	"fmt"
	"net/http"

	plaidConfig "github.com/IacopoGhilardi/personal-finance-backend/pkg/config/plaid"
	"github.com/gin-gonic/gin"
	plaid "github.com/plaid/plaid-go/plaid"
)

var plaidClient = plaidConfig.Init()
var accessToken string
var itemID string

var paymentID string

// The transfer_id is only relevant for the Transfer ACH product.
// We store the transfer_id in memory - in production, store it in a secure
// persistent data store
var transferID string

func GetAccessToken(c *gin.Context) {
	publicToken := c.PostForm("public_token")
	ctx := context.Background()

	exchangePublicTokenResp, _, err := plaidClient.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken),
	).Execute()

	fmt.Println(exchangePublicTokenResp)

	if err != nil {
		renderError(c, err)
		return
	}

	fmt.Println("ciaooooo ")

	accessToken = exchangePublicTokenResp.GetAccessToken()
	itemID = exchangePublicTokenResp.GetItemId()
	// if itemExists(strings.Split(plaidConfig.PLAID_PRODUCTS, ","), "transfer") {
	// 	transferID, err = authorizeAndCreateTransfer(ctx, plaidClient, accessToken)
	// }

	fmt.Println("public token: " + publicToken)
	fmt.Println("access token: " + accessToken)
	fmt.Println("item ID: " + itemID)

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"item_id":      itemID,
	})
}

func renderError(c *gin.Context, originalErr error) {
	if plaidError, err := plaid.ToPlaidError(originalErr); err == nil {
		// Return 200 and allow the front end to render the error.
		c.JSON(http.StatusOK, gin.H{"error": plaidError})
		return
	}

	fmt.Println("arrivo qua")

	c.JSON(http.StatusInternalServerError, gin.H{"error": originalErr.Error()})
}

func authorizeAndCreateTransfer(ctx context.Context, client *plaid.APIClient, accessToken string) (string, error) {
	accountsGetResp, _, _ := client.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*plaid.NewAccountsGetRequest(accessToken),
	).Execute()

	accountID := accountsGetResp.GetAccounts()[0].AccountId

	transferAuthorizationCreateUser := plaid.NewTransferUserInRequest("FirstName LastName")
	transferAuthorizationCreateRequest := plaid.NewTransferAuthorizationCreateRequest(
		accessToken,
		accountID,
		"credit",
		"ach",
		"1.34",
		"ppd",
		*transferAuthorizationCreateUser,
	)

	transferAuthorizationCreateResp, _, err := plaidClient.PlaidApi.TransferAuthorizationCreate(ctx).TransferAuthorizationCreateRequest(*transferAuthorizationCreateRequest).Execute()

	if err != nil {
		return "", err
	}
	authorizationID := transferAuthorizationCreateResp.GetAuthorization().Id

	// transferCreateRequest := plaid.NewTransferCreateRequest(
	// 	accessToken,
	// 	accountID,
	// 	authorizationID,
	// 	"credit",
	// 	"ach",
	// 	"1.34",
	// 	"Payment",
	// 	"ppd",
	// 	*transferAuthorizationCreateUser,
	// )
	// transferCreateResp, _, err := client.PlaidApi.TransferCreate(ctx).TransferCreateRequest(*transferCreateRequest).Execute()

	if err != nil {
		return "", err
	}

	// return transferCreateResp.GetTransfer().Id, nil
	return authorizationID, nil
}

func itemExists(array []string, product string) bool {
	for _, item := range array {
		if item == product {
			return true
		}
	}

	return false
}

package bankController

import (
	// "io/ioutil"

	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
)

func performPostJsonRequest() {
	const myUrl = "https://ob.nordigen.com/api/v2/token/new/"

	secretKey := "221b0b9e53928cbc23e91832bc85c47dcba5140a38c641eb2d97571bad57838bfa7244ccd7363d989e2c2b4b80fcded905ddd0f41ae9a3765a4bf5a229279edf"
	secretId := "3e612589-c7a2-49d8-9f1f-d24ac3488c94"

	prova := `
	{
		"secret_id": "` + secretId + `",
		"secret_key":  "` + secretKey + `"
	}`

	//json payload
	requestBody := strings.NewReader(prova)

	fmt.Println(requestBody)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Quaaaaaaaaaaaaa")
	fmt.Println(string(content))
}

func GetAccessToken(context *gin.Context) {

	url := "https://ob.nordigen.com/api/v2/token/new/"
	secretKey := "221b0b9e53928cbc23e91832bc85c47dcba5140a38c641eb2d97571bad57838bfa7244ccd7363d989e2c2b4b80fcded905ddd0f41ae9a3765a4bf5a229279edf"
	secretId := "3e612589-c7a2-49d8-9f1f-d24ac3488c94"

	jsonBody := `
	{
		"secret_id": "` + secretId + `",
		"secret_key":  "` + secretKey + `"
	}`

	requestBody := strings.NewReader(jsonBody)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	context.JSON(200, gin.H{
		"status": true,
		"data":   string(content),
	})
}

func GetBankList(context *gin.Context) {
	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNjczMzAwNzYwLCJqdGkiOiIyMDZlMGNlZWQ1NmY0YTVmOTY2NzEwNWEzNThjNzM4OCIsImlkIjoyMDY4NCwic2VjcmV0X2lkIjoiM2U2MTI1ODktYzdhMi00OWQ4LTlmMWYtZDI0YWMzNDg4Yzk0IiwiYWxsb3dlZF9jaWRycyI6WyIwLjAuMC4wLzAiLCI6Oi8wIl19.tjMT7mtqFCQQ9LwkvnUOkkIOXozVucC84d7YmzlawPY"

	url := "https://ob.nordigen.com/api/v2/institutions/"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", token)
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, true) // better way

	fmt.Printf("RESPONSE:\n%s", string(dump))

	context.JSON(200, gin.H{
		"status": true,
		"data":   string(dump),
	})

}

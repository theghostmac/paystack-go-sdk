package main

import (
	"fmt"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_transactions"
	"log"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	reqData := paystack_transactions.InitializeTransactionRequest{
		Amount: 20000,
		Email:  "customer@email.com",
	}

	initResp, err := paystack_transactions.InitializeTransaction(client, reqData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Authorization URL:", initResp.Data.AuthorizationURL)
}

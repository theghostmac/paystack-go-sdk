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

	reference := "transaction_reference"

	verifyResp, err := paystack_transactions.VerifyTransaction(client, reference)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction Status:", verifyResp.Data.Status)
}

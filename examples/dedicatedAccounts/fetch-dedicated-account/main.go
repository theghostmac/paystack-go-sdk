package main

import (
	"fmt"
	"log"

	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_accounts"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	dedicatedAccountID := 59 // Replace with a valid dedicated account ID for testing

	resp, err := paystack_accounts.FetchDedicatedAccount(client, dedicatedAccountID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Dedicated Account ID: %d\n", resp.Data.ID)
	fmt.Printf("Account Number: %s\n", resp.Data.DedicatedAccount.AccountNumber)
	fmt.Printf("Account Name: %s\n", resp.Data.DedicatedAccount.AccountName)
}

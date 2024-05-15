package main

import (
	"fmt"
	"log"

	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_accounts"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	dedicatedAccountID := 173 // Replace with a valid dedicated account ID for testing

	resp, err := paystack_accounts.DeactivateDedicatedAccount(client, dedicatedAccountID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Account ID: %d\n", resp.Data.ID)
	fmt.Printf("Account Number: %s\n", resp.Data.AccountNumber)
	fmt.Printf("Account Name: %s\n", resp.Data.AccountName)
}

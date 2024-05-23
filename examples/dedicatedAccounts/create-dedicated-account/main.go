package main

import (
	"log"

	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_accounts"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"
)

func main() {
	client, err := paystack_client.NewClient("YOUR_SECRET_KEY")
	if err != nil {
		log.Fatalf("Failed to create Paystack client: %v", err)
	}

	reqData := paystack_accounts.CreateDedicatedAccountRequest{
		Customer:      "CUS_xxxxxxxxx", // Replace with a valid customer ID or code
		PreferredBank: "wema-bank",
	}

	resp, err := paystack_accounts.CreateDedicatedAccount(client, reqData)
	if err != nil {
		log.Fatalf("Failed to create dedicated account: %v", err)
	}

	log.Printf("Create Dedicated Account response: %v", resp)
	log.Printf("Account ID: %d", resp.Data.ID)
	log.Printf("Account Number: %s", resp.Data.AccountNumber)
	log.Printf("Account Name: %s", resp.Data.AccountName)
}

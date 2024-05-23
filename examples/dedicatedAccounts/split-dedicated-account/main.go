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

	reqData := paystack_accounts.SplitDedicatedAccountRequest{
		Customer:      "CUS_aso0xkdnjrfhrgu", // Replace with a valid customer ID or code
		PreferredBank: "wema-bank",
		SplitCode:     "SPL_e7jnRLtzla", // Replace with a valid split code
	}

	resp, err := paystack_accounts.SplitDedicatedAccount(client, reqData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Account ID: %d\n", resp.Data.ID)
	fmt.Printf("Account Number: %s\n", resp.Data.AccountNumber)
	fmt.Printf("Account Name: %s\n", resp.Data.AccountName)
}

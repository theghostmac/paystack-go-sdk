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

	resp, err := paystack_accounts.FetchBankProviders(client)
	if err != nil {
		log.Fatal(err)
	}

	for _, provider := range resp.Data {
		fmt.Printf("Bank Provider: %v\n", provider)
	}
}

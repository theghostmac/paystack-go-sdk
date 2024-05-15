package main

import (
	"fmt"
	"log"

	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_transfers"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	queryParams := map[string]string{
		"perPage": "50",
		"page":    "1",
	}

	resp, err := paystack_transfers.ListTransfers(client, queryParams)
	if err != nil {
		log.Fatal(err)
	}

	for _, transfer := range resp.Data {
		fmt.Printf("Transfer: %v\n", transfer)
	}
}

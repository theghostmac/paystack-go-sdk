package main

import (
	"fmt"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_transactions"
	"log"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	idOrReference := "valid_transaction_reference" // Replace with a valid transaction reference or ID

	resp, err := paystack_transactions.ViewTransactionTimeline(client, idOrReference)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction Timeline: %+v\n", resp.Data.History)
}

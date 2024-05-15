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

	queryParams := map[string]string{
		"perPage": "2",
		"page":    "1",
	}

	listResp, err := paystack_transactions.ListTransactions(client, queryParams)
	if err != nil {
		log.Fatal(err)
	}
	for _, transaction := range listResp.Data {
		fmt.Printf("Transaction ID: %d, Status: %s, Amount: %d\n", transaction.ID, transaction.Status, transaction.Amount)
	}
}

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

	transactionID := uint64(3795390315) // Example transaction ID

	fetchResp, err := paystack_transactions.FetchTransaction(client, transactionID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction ID: %d, Status: %s, Amount: %d\n", fetchResp.Data.ID, fetchResp.Data.Status, fetchResp.Data.Amount)
}

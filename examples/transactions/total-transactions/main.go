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

	queryParams := map[string]string{
		"perPage": "50",
		"page":    "1",
	}

	resp, err := paystack_transactions.GetTransactionTotals(client, queryParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total Transactions: %d\n", resp.Data.TotalTransactions)
	fmt.Printf("Unique Customers: %d\n", resp.Data.UniqueCustomers)
	fmt.Printf("Total Volume: %d\n", resp.Data.TotalVolume)
}

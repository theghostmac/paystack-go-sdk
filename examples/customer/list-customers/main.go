package main

import (
	"fmt"
	"log"

	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmac/paystack-go-sdk/pkg/paystack_customer"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	queryParams := map[string]string{
		"perPage": "3",
		"page":    "1",
	}

	resp, err := paystack_customer.ListCustomers(client, queryParams)
	if err != nil {
		log.Fatal(err)
	}
	for _, customer := range resp.Data {
		fmt.Printf("Customer: %v\n", customer)
	}
}

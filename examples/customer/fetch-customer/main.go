package main

import (
	"fmt"
	"log"

	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_client"
	"github.com/theghostmc/paystack-go-sdk/pkg/paystack_customer"
)

func main() {
	apiKey := "your_api_key"
	client, err := paystack_client.NewClient(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	emailOrCode := "CUS_aso0xkdnjrfhrgu" // Use a valid customer code

	resp, err := paystack_customer.FetchCustomer(client, emailOrCode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Customer ID: %d\n", resp.Data.ID)
	fmt.Printf("Customer Email: %s\n", resp.Data.Email)
	fmt.Printf("Customer Code: %s\n", resp.Data.CustomerCode)
}

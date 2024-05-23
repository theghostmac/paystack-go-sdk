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

	customerCode := "CUS_aso0xkdnjrfhrgu" // Use a valid customer code
	reqData := paystack_customer.UpdateCustomerRequest{
		FirstName: "BoJack",
		LastName:  "Horseman",
	}

	resp, err := paystack_customer.UpdateCustomer(client, customerCode, reqData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Customer ID: %d\n", resp.Data.ID)
	fmt.Printf("Customer First Name: %s\n", resp.Data.FirstName)
	fmt.Printf("Customer Last Name: %s\n", resp.Data.LastName)
}
